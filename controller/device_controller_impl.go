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
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/pakaiwa/api/model/api"
	"github.com/pakaiwa/api/service"
	"github.com/pakaiwa/api/utils"
	"net/http"
)

type DeviceControllerImpl struct {
	DeviceService service.DeviceService
}

func NewDeviceController(deviceService service.DeviceService) DeviceController {
	return &DeviceControllerImpl{DeviceService: deviceService}
}

func (controller *DeviceControllerImpl) RegisterRoutes(router *httprouter.Router) {
	router.POST("/devices", controller.AddDevice)
	router.GET("/devices", controller.GetAllDevices)
	router.GET("/devices/:deviceId", controller.GetDeviceById)
	router.DELETE("/devices/:deviceId", controller.DeleteDevice)
}

func (controller *DeviceControllerImpl) AddDevice(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Println("Invoke AddDevice Controller")
	req := api.DeviceAddRq{}
	api.ReadFromRequestBody(request, &req)

	res := controller.DeviceService.AddDevice(request.Context(), req)
	webResponse := api.ResponseAPI{
		Code:   http.StatusCreated,
		Status: "OK",
		Data:   res,
		Meta: &api.Meta{
			Location: utils.GetMetaLocation(request) + res.Id,
		},
	}

	api.WriteToResponseBody(writer, webResponse.Code, webResponse)
}

func (controller *DeviceControllerImpl) DeleteDevice(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Println("Invoke DeleteDevice Controller")

	deviceId := params.ByName("deviceId")
	controller.DeviceService.DeleteDevice(request.Context(), deviceId)

	writer.WriteHeader(http.StatusNoContent)
}

func (controller *DeviceControllerImpl) GetDeviceById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Println("Invoke GetDeviceById Controller")
	deviceId := params.ByName("deviceId")
	res := controller.DeviceService.GetDevice(request.Context(), deviceId)

	webResponse := api.ResponseAPI{
		Code:   200,
		Status: "OK",
		Data:   res,
		Meta:   nil,
	}

	api.WriteToResponseBody(writer, webResponse.Code, webResponse)

}

func (controller *DeviceControllerImpl) GetAllDevices(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Println("Invoke GetAllDevices Controller")

	res := controller.DeviceService.GetAllDevices(request.Context())

	webResponse := api.ResponseAPI{
		Code:   200,
		Status: "OK",
		Data:   res,
		Meta:   nil,
	}

	api.WriteToResponseBody(writer, webResponse.Code, webResponse)

}
