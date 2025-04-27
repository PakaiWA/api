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
	"fmt"
	"github.com/google/uuid"
	"github.com/pakaiwa/api/helper"
	"github.com/pakaiwa/api/model/entity"
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
	//TODO implement me
	panic("implement me")
}

func (repository *DeviceRepositoryImpl) FindDeviceById(ctx context.Context, tx *sql.Tx, categoryId int) (entity.Device, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *DeviceRepositoryImpl) FindAllDevice(ctx context.Context, tx *sql.Tx) []entity.Device {
	//TODO implement me
	panic("implement me")
}
