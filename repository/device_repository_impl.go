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
	"github.com/pakaiwa/api/model/entity"
)

type DeviceRepositoryImpl struct{}

func NewDeviceRepository() DeviceRepository {
	return &DeviceRepositoryImpl{}
}

func (device DeviceRepositoryImpl) AddDevice(ctx context.Context, tx *sql.Tx, category entity.Device) entity.Device {
	//TODO implement me
	panic("implement me")
}

func (device DeviceRepositoryImpl) DeleteDevice(ctx context.Context, tx *sql.Tx, category entity.Device) {
	//TODO implement me
	panic("implement me")
}

func (device DeviceRepositoryImpl) FindDeviceById(ctx context.Context, tx *sql.Tx, categoryId int) (entity.Device, error) {
	//TODO implement me
	panic("implement me")
}

func (device DeviceRepositoryImpl) FindAllDevice(ctx context.Context, tx *sql.Tx) []entity.Device {
	//TODO implement me
	panic("implement me")
}
