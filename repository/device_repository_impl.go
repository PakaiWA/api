// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sun 27/04/25 18.09
// @project api repository
//

package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/pakaiwa/api/config"
	"github.com/pakaiwa/api/helper"
	"github.com/pakaiwa/api/model/entity"
	"github.com/pakaiwa/api/utils"
)

type DeviceRepositoryImpl struct{}

func NewDeviceRepository() DeviceRepository {
	return &DeviceRepositoryImpl{}
}

func (repository *DeviceRepositoryImpl) AddDevice(ctx context.Context, tx *sql.Tx, device entity.Device) (entity.Device, error) {
	fmt.Println("Invoke AddDevice Repository")

	var count int
	CountDeviceSQL := config.GetCountDeviceSQL()
	err := tx.QueryRowContext(ctx, CountDeviceSQL, device.Name, ctx.Value("userEmail").(string)).Scan(&count)
	if err != nil {
		return device, err
	}

	if count != 0 {
		return device, errors.New("error: Device id already exists, choose another one")
	}

	deviceId := utils.GenerateUUID()
	AddDeviceSQL := config.GetAddDeviceSQL()
	fmt.Println(AddDeviceSQL, deviceId, ctx.Value("userEmail").(string), device.Name)

	err = tx.QueryRowContext(ctx, AddDeviceSQL, deviceId, ctx.Value("userEmail").(string), device.Name).
		Scan(&device.Name, &device.Status, &device.CreatedAt)
	if err != nil {
		return device, err
	}

	fmt.Println("Success insert device", device)
	return device, nil
}

func (repository *DeviceRepositoryImpl) DeleteDevice(ctx context.Context, tx *sql.Tx, device entity.Device) {
	fmt.Println("Invoke DeleteDevice Repository")

	SQL := config.GetDeleteDeviceSQL()
	_, err := tx.ExecContext(ctx, SQL, device.Name, ctx.Value("userEmail").(string))
	helper.PanicIfError(err)
}

func (repository *DeviceRepositoryImpl) FindDeviceById(ctx context.Context, tx *sql.Tx, deviceId string) (entity.Device, error) {
	fmt.Println("Invoke FindDeviceById Repository")

	SQL := config.GetDeviceByIdSQL()
	rows, err := tx.QueryContext(ctx, SQL, deviceId, ctx.Value("userEmail").(string))
	helper.PanicIfError(err)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			helper.PanicIfError(err)
		}
	}(rows)

	var connectedAt sql.NullString
	var disconnectedAt sql.NullString
	var disconnectedReason sql.NullString

	device := entity.Device{}

	if rows.Next() {
		err := rows.Scan(
			&device.Name,
			&device.Status,
			&device.PhoneNumber,
			&device.CreatedAt,
			&connectedAt,
			&disconnectedAt,
			&disconnectedReason,
		)
		helper.PanicIfError(err)

		device.ConnectedAt = utils.SafeString(connectedAt)
		device.DisconnectedAt = utils.SafeString(disconnectedAt)
		device.DisconnectedReason = utils.SafeString(disconnectedReason)

		return device, nil
	} else {
		return device, errors.New("error: Device not found")
	}
}

func (repository *DeviceRepositoryImpl) GetAllDevices(ctx context.Context, tx *sql.Tx) []entity.Device {
	fmt.Println("Invoke GetAllDevices Repository")

	SQL := config.GetAllDevicesSQL()
	rows, err := tx.QueryContext(ctx, SQL, ctx.Value("userEmail").(string))
	helper.PanicIfError(err)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			helper.PanicIfError(err)
		}
	}(rows)

	var connectedAt sql.NullString
	var disconnectedAt sql.NullString
	var disconnectedReason sql.NullString

	var devices []entity.Device

	for rows.Next() {
		device := entity.Device{}
		err := rows.Scan(
			&device.Name,
			&device.Status,
			&device.PhoneNumber,
			&device.CreatedAt,
			&connectedAt,
			&disconnectedAt,
			&disconnectedReason,
		)
		helper.PanicIfError(err)
		device.ConnectedAt = utils.SafeString(connectedAt)
		device.DisconnectedAt = utils.SafeString(disconnectedAt)
		device.DisconnectedReason = utils.SafeString(disconnectedReason)

		devices = append(devices, device)
	}

	return devices
}
