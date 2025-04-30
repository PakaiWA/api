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
	"github.com/google/uuid"
	"github.com/pakaiwa/api/helper"
	"github.com/pakaiwa/api/model/entity"
	"github.com/pakaiwa/api/utils"
)

type DeviceRepositoryImpl struct{}

func NewDeviceRepository() DeviceRepository {
	return &DeviceRepositoryImpl{}
}

func (repository *DeviceRepositoryImpl) AddDevice(ctx context.Context, tx *sql.Tx, device entity.Device) entity.Device {
	fmt.Println("Invoke AddDevice Repository")

	deviceId := uuid.New()

	SQL := "insert into device.user_devices (uuid, name) values ($1, $2) RETURNING name, status, created_at"

	fmt.Println(SQL, deviceId, device.Name)

	err := tx.QueryRowContext(ctx, SQL, deviceId, device.Name).
		Scan(&device.Name, &device.Status, &device.CreatedAt)

	helper.PanicIfError(err)
	fmt.Println("Success insert device", device)

	return device
}

func (repository *DeviceRepositoryImpl) DeleteDevice(ctx context.Context, tx *sql.Tx, device entity.Device) {
	fmt.Println("Invoke DeleteDevice Repository")

	SQL := "delete from device.user_devices where name = $1"
	_, err := tx.ExecContext(ctx, SQL, device.Name)
	helper.PanicIfError(err)
}

func (repository *DeviceRepositoryImpl) FindDeviceById(ctx context.Context, tx *sql.Tx, deviceId string) (entity.Device, error) {
	fmt.Println("Invoke FindDeviceById Repository")

	SQL := "select name, status, phone_number, created_at, connected_at, disconnected_at, disconnected_reason from device.user_devices where name = $1"
	rows, err := tx.QueryContext(ctx, SQL, deviceId)
	helper.PanicIfError(err)
	defer rows.Close()

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

	SQL := "select name, status, phone_number, created_at, connected_at, disconnected_at, disconnected_reason from device.user_devices"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

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
