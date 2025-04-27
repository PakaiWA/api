// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sun 27/04/25 17.12
// @project api app
//

package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/pakaiwa/api/controller"
	"github.com/pakaiwa/api/exception"
)

func NewRouter(deviceController controller.DeviceController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/devices", deviceController.AddDevice)
	router.GET("/devices", deviceController.GetAllDevices)
	router.GET("/devices/:categoryId", deviceController.GetDeviceById)
	router.DELETE("/devices/:categoryId", deviceController.DeleteDevice)

	router.PanicHandler = exception.ErrorHandler

	return router
}
