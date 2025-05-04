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
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pakaiwa/api/config"
	"github.com/pakaiwa/api/exception"
	"github.com/pakaiwa/api/helper"
	"github.com/pakaiwa/api/model/api"
	"github.com/pakaiwa/api/repository"
	"github.com/pakaiwa/pakaiwa"
	"github.com/pakaiwa/pakaiwa/store/sqlstore"
	waLog "github.com/pakaiwa/pakaiwa/util/log"
	"log"
	"sync"
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

var (
	clientMap  = make(map[string]*pakaiwa.Client)
	clientLock = sync.Mutex{}
)

func (service QRServiceImpl) GetQRCode(ctx context.Context, deviceId string) api.QRCodeRs {
	fmt.Println("Invoke GetQRCode Service")
	tx, conn, err := helper.DBTransaction(ctx, service.DB)
	helper.PanicIfError(err)
	defer conn.Release()
	defer helper.CommitOrRollback(ctx, tx)

	device, err := service.DeviceRepository.FindDeviceById(ctx, tx, deviceId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	fmt.Println("Get QR Code Success", device)

	clientLock.Lock()
	defer clientLock.Unlock()

	if existingClient, ok := clientMap[device.Id]; ok {
		if existingClient.IsConnected() {
			log.Println("Client already connected for device:", device.Id)
		}
	}

	// Load device and store
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	container, err := sqlstore.New(config.GetDBCon(), dbLog)
	if err != nil {
		panic(err)
	}
	store := container.NewDevice()
	client := pakaiwa.NewClient(store, nil)

	clientMap[device.Id] = client

	go client.Connect() // non-blocking connect

	qrCode := api.QRCodeRs{}
	qrCode.QRCode = "https://api.pakaiwa.com/qr/" + device.Id
	qrCode.ImageUrl = "https://api.pakaiwa.com/qr/" + device.Name
	return qrCode
}
