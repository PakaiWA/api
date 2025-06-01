// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sun 27/04/25 18.19
// @project api https://github.com/PakaiWA/api/tree/main/service
//

package service

import (
	"context"
	"github.com/pakaiwa/api/logx"
	"net/http"

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

func (service *DeviceServiceImpl) GetDeviceById(ctx context.Context, id string) (api.DeviceRs, error) {
	logx.DebugCtx(ctx, "Invoke GetDeviceById Service")

	tx, conn, err := helper.DBTransaction(ctx, service.DB)
	logx.DebugCtx(ctx, "Invoke DBTransaction")
	helper.PanicIfError(err)
	defer conn.Release()
	defer helper.CommitOrRollback(ctx, tx)

	device, err := service.DeviceRepository.FindDeviceById(ctx, tx, id)
	if err != nil {
		return api.DeviceRs{}, err
	}

	return api.ToDeviceResponse(ctx, device), nil
}

func (service *DeviceServiceImpl) DeleteDevice(ctx context.Context, id string) {
	logx.DebugCtx(ctx, "Invoke DeleteDevice Service")

	tx, conn, err := helper.DBTransaction(ctx, service.DB)
	logx.DebugCtx(ctx, "Invoke DBTransaction")
	helper.PanicIfError(err)
	defer conn.Release()
	defer helper.CommitOrRollback(ctx, tx)

	device, err := service.DeviceRepository.FindDeviceById(ctx, tx, id)
	if err != nil {
		panic(exception.NewHTTPError(http.StatusNotFound, err.Error()))
	}

	service.DeviceRepository.DeleteDevice(ctx, tx, device)
}

func (service *DeviceServiceImpl) GetAllDevices(ctx context.Context) []api.DeviceRs {
	logx.DebugCtx(ctx, "Invoke GetAllDevices Service")

	tx, conn, err := helper.DBTransaction(ctx, service.DB)
	logx.DebugCtx(ctx, "Invoke DBTransaction")
	helper.PanicIfError(err)
	defer conn.Release()
	defer helper.CommitOrRollback(ctx, tx)

	devices := service.DeviceRepository.GetAllDevices(ctx, tx)

	return api.ToDeviceResponses(ctx, devices)
}

func (service *DeviceServiceImpl) GetDevice(ctx context.Context, id string) api.DeviceRs {
	logx.DebugCtx(ctx, "Invoke GetDevice Service")

	tx, conn, err := helper.DBTransaction(ctx, service.DB)
	logx.DebugCtx(ctx, "Invoke DBTransaction")
	helper.PanicIfError(err)
	defer conn.Release()
	defer helper.CommitOrRollback(ctx, tx)

	device, err := service.DeviceRepository.FindDeviceById(ctx, tx, id)
	if err != nil {
		panic(exception.NewHTTPError(http.StatusNotFound, err.Error()))
	}

	return api.ToDeviceResponse(ctx, device)
}

func (service *DeviceServiceImpl) AddDevice(ctx context.Context, req api.DeviceAddRq) api.DeviceRs {
	logx.DebugCtx(ctx, "Invoke AddDevice Service")
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)

	tx, conn, err := helper.DBTransaction(ctx, service.DB)
	logx.DebugCtx(ctx, "Invoke DBTransaction")
	helper.PanicIfError(err)
	defer conn.Release()
	defer helper.CommitOrRollback(ctx, tx)

	device := entity.Device{
		Name: req.DeviceId,
	}

	device, err = service.DeviceRepository.AddDevice(ctx, tx, device)
	if err != nil {
		panic(exception.NewHTTPError(http.StatusBadRequest, err.Error()))
	}

	return api.ToDeviceResponse(ctx, device)
}
