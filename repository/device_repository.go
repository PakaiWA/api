// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sun 27/04/25 18.06
// @project api repository
//

package repository

import (
	"context"
	"database/sql"
	"github.com/pakaiwa/api/model/entity"
)

type DeviceRepository interface {
	AddDevice(ctx context.Context, tx *sql.Tx, device entity.Device) entity.Device
	DeleteDevice(ctx context.Context, tx *sql.Tx, device entity.Device)
	FindDeviceById(ctx context.Context, tx *sql.Tx, deviceId string) (entity.Device, error)
	GetAllDevices(ctx context.Context, tx *sql.Tx) []entity.Device
}
