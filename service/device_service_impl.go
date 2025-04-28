// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sun 27/04/25 18.19
// @project api service
//

package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/pakaiwa/api/exception"
	"github.com/pakaiwa/api/helper"
	"github.com/pakaiwa/api/model/api"
	"github.com/pakaiwa/api/model/entity"
	"github.com/pakaiwa/api/repository"
)

type DeviceServiceImpl struct {
	DeviceRepository repository.DeviceRepository
	DB               *sql.DB
	Validate         *validator.Validate
}

func NewDeviceService(deviceRepository repository.DeviceRepository, DB *sql.DB, validate *validator.Validate) DeviceService {
	return &DeviceServiceImpl{
		DeviceRepository: deviceRepository,
		DB:               DB,
		Validate:         validate,
	}
}

func (service *DeviceServiceImpl) DeleteDevice(ctx context.Context, id string) {
	//TODO implement me
	panic("implement me")
}

func (service *DeviceServiceImpl) GetAllDevices(ctx context.Context) []api.DeviceRs {
	//TODO implement me
	panic("implement me")
}

func (service *DeviceServiceImpl) GetDevice(ctx context.Context, id string) api.DeviceRs {
	fmt.Println("Invoke GetDevice Service")

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	devices, err := service.DeviceRepository.FindDeviceById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return api.ToDeviceResponse(devices)
}

func (service *DeviceServiceImpl) AddDevice(ctx context.Context, req api.DeviceAddRq) api.DeviceRs {
	fmt.Println("Invoke AddDevice Service")
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	device := entity.Device{
		Name: req.DeviceId,
	}

	device = service.DeviceRepository.AddDevice(ctx, tx, device)

	return api.ToDeviceResponse(device)
}
