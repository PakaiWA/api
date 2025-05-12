// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sat 03/05/25 15.54
// @project api https://github.com/PakaiWA/api/tree/main/service
//

package service

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pakaiwa/api/repository"
	"github.com/pakaiwa/api/session"
)

type QRServiceImpl struct {
	DeviceRepository repository.DeviceRepository
	DB               *pgxpool.Pool
}

func NewQRService(deviceRepository repository.DeviceRepository, DB *pgxpool.Pool) *QRServiceImpl {
	return &QRServiceImpl{
		DeviceRepository: deviceRepository,
		DB:               DB,
	}
}

func (service QRServiceImpl) GetQRCode(ctx context.Context, deviceId string) string {
	fmt.Println("Invoke GetQRCode Service")

	pakaiwaClient := session.NewDevicePakaiWA(deviceId)

	qrCode := session.QRHandler(ctx, pakaiwaClient)

	return qrCode
}
