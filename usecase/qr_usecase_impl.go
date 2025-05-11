// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sun 04/05/25 22.05
// @project api https://github.com/PakaiWA/api/tree/main/usecase
//

package usecase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pakaiwa/api/exception"
	"github.com/pakaiwa/api/model/api"
	"github.com/pakaiwa/api/service"
	"github.com/pakaiwa/api/session"
	"github.com/pakaiwa/api/utils"
)

type QRUsecaseImpl struct {
	QRService     service.QRService
	DeviceService service.DeviceService
}

func NewQRUsecase(QRService service.QRService, DeviceService service.DeviceService) QRUsecase {
	return &QRUsecaseImpl{QRService: QRService, DeviceService: DeviceService}
}

func (usecase QRUsecaseImpl) GetQRCode(ctx context.Context, request *http.Request) api.QRCodeRs {
	fmt.Println("Invoke GetQRCode Usecase")
	// Check if the device is registered
	device, err := usecase.DeviceService.GetDeviceById(ctx, request.URL.Query().Get("device_id"))
	if err != nil {
		panic(exception.NewHTTPError(http.StatusNotFound, err.Error()))
	}
	fmt.Println(device)

	pakaiwaClient := session.NewDevicePakaiWA(device.Id)

	//QR service
	qrCode := session.QRHandler(pakaiwaClient)

	QRResponse := api.QRCodeRs{
		QRCode:   qrCode,
		ImageUrl: utils.GetHost(request) + "/qr/show?qrCode=" + qrCode,
	}

	return QRResponse
}
