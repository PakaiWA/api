// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sun 27/04/25 18.09
// @project api https://github.com/PakaiWA/api/tree/main/repository
//

package repository

import (
	"context"
	"errors"
	"github.com/pakaiwa/api/logx"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/pakaiwa/api/config"
	"github.com/pakaiwa/api/helper"
	"github.com/pakaiwa/api/model/entity"
	"github.com/pakaiwa/api/utils"
)

type DeviceRepositoryImpl struct{}

func NewDeviceRepository() DeviceRepository {
	return &DeviceRepositoryImpl{}
}

func (repository *DeviceRepositoryImpl) AddDevice(ctx context.Context, tx pgx.Tx, device entity.Device) (entity.Device, error) {
	logx.DebugCtx(ctx, "Invoke AddDevice Repository")

	var count int
	CountDeviceSQL := config.GetCountDeviceSQL()
	err := tx.QueryRow(ctx, CountDeviceSQL, strings.ToLower(device.Name), ctx.Value("userEmail").(string)).Scan(&count)
	if err != nil {
		return device, err
	}

	if count != 0 {
		return device, errors.New("error: Device id already exists, choose another one")
	}

	deviceId := utils.GenerateUUID()
	AddDeviceSQL := config.GetAddDeviceSQL()
	logx.DebugfCtx(ctx, AddDeviceSQL, deviceId, ctx.Value("userEmail").(string), strings.ToLower(device.Name))

	err = tx.QueryRow(ctx, AddDeviceSQL, deviceId, ctx.Value("userEmail").(string), strings.ToLower(device.Name)).
		Scan(&device.Name, &device.Status, &device.CreatedAt)
	if err != nil {
		return device, err
	}

	logx.DebugfCtx(ctx, "Success insert device", device)
	return device, nil
}

func (repository *DeviceRepositoryImpl) DeleteDevice(ctx context.Context, tx pgx.Tx, device entity.Device) {
	logx.DebugCtx(ctx, "Invoke DeleteDevice Repository")

	SQL := config.GetDeleteDeviceSQL()
	_, err := tx.Exec(ctx, SQL, strings.ToLower(device.Name), ctx.Value("userEmail").(string))
	helper.PanicIfError(err)
}

func (repository *DeviceRepositoryImpl) FindDeviceById(ctx context.Context, tx pgx.Tx, deviceId string) (entity.Device, error) {
	logx.DebugCtx(ctx, "Invoke FindDeviceById Repository")

	SQL := config.GetDeviceByIdSQL()
	rows, err := tx.Query(ctx, SQL, strings.ToLower(deviceId), ctx.Value("userEmail").(string))
	helper.PanicIfError(err)
	defer rows.Close()

	device := entity.Device{}

	if rows.Next() {
		err := rows.Scan(
			&device.Id,
			&device.Name,
			&device.Status,
			&device.PhoneNumber,
			&device.CreatedAt,
			&device.ConnectedAt,
			&device.DisconnectedAt,
			&device.DisconnectedReason,
		)
		helper.PanicIfError(err)
		return device, nil
	} else {
		return device, errors.New("error: Device not found")
	}
}

func (repository *DeviceRepositoryImpl) GetAllDevices(ctx context.Context, tx pgx.Tx) []entity.Device {
	logx.DebugCtx(ctx, "Invoke GetAllDevices Repository")

	SQL := config.GetAllDevicesSQL()
	logx.DebugfCtx(ctx, SQL, ctx.Value("userEmail").(string))
	rows, err := tx.Query(ctx, SQL, ctx.Value("userEmail").(string))
	helper.PanicIfError(err)
	defer rows.Close()

	var devices []entity.Device

	for rows.Next() {
		device := entity.Device{}
		err := rows.Scan(
			&device.Name,
			&device.Status,
			&device.PhoneNumber,
			&device.CreatedAt,
			&device.ConnectedAt,
			&device.DisconnectedAt,
			&device.DisconnectedReason,
		)
		helper.PanicIfError(err)
		devices = append(devices, device)
	}

	return devices
}
