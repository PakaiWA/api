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
	"fmt"
	"github.com/KAnggara75/scc2go"
	"github.com/go-playground/validator/v10"
	_ "github.com/jackc/pgx/v5"
	"github.com/pakaiwa/api/app"
	"github.com/pakaiwa/api/controller"
	"github.com/pakaiwa/api/helper"
	"github.com/pakaiwa/api/repository"
	"github.com/pakaiwa/api/service"
	"net/http"
	"os"
)

func init() {
	scc2go.GetEnv(os.Getenv("SCC_URL"), os.Getenv("AUTH"))
}

func main() {

	db := app.NewDBConn()
	validate := validator.New()
	deviceRepository := repository.NewDeviceRepository()
	deviceService := service.NewDeviceService(deviceRepository, db, validate)
	deviceController := controller.NewDeviceController(deviceService)
	router := app.NewRouter(deviceController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
		//Handler: middleware.NewAuthMiddleware(router),
	}

	fmt.Println("Listening on port 3000")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
