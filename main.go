// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sun 27/04/25 17.10
// @project api main
//

package main

import (
	"github.com/pakaiwa/api/app"
	"github.com/pakaiwa/api/controller"
	"github.com/pakaiwa/api/helper"
	"github.com/pakaiwa/api/service"
	"net/http"
)

func main() {
	categoryService := service.NewDeviceService()
	categoryController := controller.NewDeviceController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
		//Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
