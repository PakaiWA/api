// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sat 03/05/25 15.54
// @project api service
//

package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pakaiwa/api/exception"
	"github.com/pakaiwa/api/helper"
	"github.com/pakaiwa/api/model/api"
	"github.com/pakaiwa/api/repository"
)

type QRServiceImpl struct {
	DeviceRepository repository.DeviceRepository
	DB               *sql.DB
}

func NewQRService(deviceRepository repository.DeviceRepository, DB *sql.DB) *QRServiceImpl {
	return &QRServiceImpl{
		DeviceRepository: deviceRepository,
		DB:               DB,
	}
}

func (service QRServiceImpl) GetQRCode(ctx context.Context, deviceId string) api.QRCodeRs {
	fmt.Println("Invoke GetQRCode Service")
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	device, err := service.DeviceRepository.FindDeviceById(ctx, tx, deviceId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	fmt.Println("Get QR Code Success", device)

	qrCode := api.QRCodeRs{}
	qrCode.QRCode = "https://api.pakaiwa.com/qr/" + device.Id
	qrCode.ImageUrl = "https://api.pakaiwa.com/qr/" + device.Name
	return qrCode
}
