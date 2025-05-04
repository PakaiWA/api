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
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pakaiwa/api/exception"
	"github.com/pakaiwa/api/helper"
	"github.com/pakaiwa/api/model/api"
	"github.com/pakaiwa/api/model/entity"
	"github.com/pakaiwa/api/repository"
)

type DeviceServiceImpl struct {
	DeviceRepository repository.DeviceRepository
	DB               *pgxpool.Pool
	Validate         *validator.Validate
}

func NewDeviceService(deviceRepository repository.DeviceRepository, DB *pgxpool.Pool, validate *validator.Validate) DeviceService {
	return &DeviceServiceImpl{
		DeviceRepository: deviceRepository,
		DB:               DB,
		Validate:         validate,
	}
}

func (service *DeviceServiceImpl) DeleteDevice(ctx context.Context, id string) {
	fmt.Println("Invoke DeleteDevice Service")

	tx, conn, err := helper.DBTransaction(ctx, service.DB)
	helper.PanicIfError(err)
	defer conn.Release()
	defer helper.CommitOrRollback(ctx, tx)

	device, err := service.DeviceRepository.FindDeviceById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.DeviceRepository.DeleteDevice(ctx, tx, device)
}

func (service *DeviceServiceImpl) GetAllDevices(ctx context.Context) []api.DeviceRs {
	fmt.Println("Invoke GetAllDevices Service")

	tx, conn, err := helper.DBTransaction(ctx, service.DB)
	helper.PanicIfError(err)
	defer conn.Release()
	defer helper.CommitOrRollback(ctx, tx)

	devices := service.DeviceRepository.GetAllDevices(ctx, tx)

	return api.ToDeviceResponses(devices)
}

func (service *DeviceServiceImpl) GetDevice(ctx context.Context, id string) api.DeviceRs {
	fmt.Println("Invoke GetDevice Service")

	tx, conn, err := helper.DBTransaction(ctx, service.DB)
	helper.PanicIfError(err)
	defer conn.Release()
	defer helper.CommitOrRollback(ctx, tx)

	device, err := service.DeviceRepository.FindDeviceById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return api.ToDeviceResponse(device)
}

func (service *DeviceServiceImpl) AddDevice(ctx context.Context, req api.DeviceAddRq) api.DeviceRs {
	fmt.Println("Invoke AddDevice Service")
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)

	tx, conn, err := helper.DBTransaction(ctx, service.DB)
	helper.PanicIfError(err)
	defer conn.Release()
	defer helper.CommitOrRollback(ctx, tx)

	device := entity.Device{
		Name: req.DeviceId,
	}

	device, err = service.DeviceRepository.AddDevice(ctx, tx, device)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	return api.ToDeviceResponse(device)
}
