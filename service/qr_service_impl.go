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
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pakaiwa/api/exception"
	"github.com/pakaiwa/api/helper"
	"github.com/pakaiwa/api/model/api"
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

func (service QRServiceImpl) GetQRCode(ctx context.Context, deviceId string) api.QRCodeRs {
	fmt.Println("Invoke GetQRCode Service")
	tx, conn, err := helper.DBTransaction(ctx, service.DB)
	helper.PanicIfError(err)
	defer conn.Release()
	defer helper.CommitOrRollback(ctx, tx)

	device, err := service.DeviceRepository.FindDeviceById(ctx, tx, deviceId)
	if err != nil {
		panic(exception.NewHTTPError(http.StatusNotFound, err.Error()))
	}

	pakaiwaClient := session.NewDevicePakaiWA(device.Id)

	qrCode := session.QRHandler(pakaiwaClient)

	QRResponse := api.QRCodeRs{
		QRCode: qrCode,
	}
	return QRResponse
}
