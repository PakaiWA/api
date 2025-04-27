// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sun 27/04/25 18.20
// @project api controller
//

package controller

import (
	"github.com/julienschmidt/httprouter"
	"github.com/pakaiwa/api/service"
	"net/http"
)

type DeviceControllerImpl struct {
	DeviceService service.DeviceService
}

func NewDeviceController(deviceService service.DeviceService) DeviceController {
	return &DeviceControllerImpl{DeviceService: deviceService}
}

func (device DeviceControllerImpl) AddDevice(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

func (device DeviceControllerImpl) DeleteDevice(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

func (device DeviceControllerImpl) GetDeviceById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

func (device DeviceControllerImpl) GetAllDevices(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}
