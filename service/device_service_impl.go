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
	"github.com/pakaiwa/api/model/api"
)

type DeviceServiceImpl struct{}

func (device DeviceServiceImpl) DeleteDevice(ctx context.Context, id string) {
	//TODO implement me
	panic("implement me")
}

func (device DeviceServiceImpl) GetAllDevices(ctx context.Context) []api.DeviceRs {
	//TODO implement me
	panic("implement me")
}

func (device DeviceServiceImpl) GetDevice(ctx context.Context, id string) api.DeviceRs {
	//TODO implement me
	panic("implement me")
}

func (device DeviceServiceImpl) AddDevice(ctx context.Context, req api.DeviceAddRq) api.DeviceRs {
	//TODO implement me
	panic("implement me")
}

func NewDeviceService() DeviceService {
	return &DeviceServiceImpl{}
}
