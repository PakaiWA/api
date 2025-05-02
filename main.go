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
	"github.com/julienschmidt/httprouter"
	"github.com/pakaiwa/api/app"
	"github.com/pakaiwa/api/controller"
	"github.com/pakaiwa/api/exception"
	"github.com/pakaiwa/api/helper"
	"github.com/pakaiwa/api/repository"
	"github.com/pakaiwa/api/service"
	"net/http"
	"os"
)

func init() {
	scc2go.GetEnv(os.Getenv("SCC_URL"), os.Getenv("AUTH"))
	app.NewRedisClient()
}

func main() {
	db := app.NewDBConn()
	router := httprouter.New()
	validate := validator.New()

	deviceController := controller.NewDeviceController(service.NewDeviceService(repository.NewDeviceRepository(), db, validate))
	deviceController.RegisterRoutes(router)

	userController := controller.NewUserController(service.NewUserService(repository.NewUserRepo(), db, validate))
	userController.RegisterRoutes(router)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	router.PanicHandler = exception.ErrorHandler
	fmt.Println("Listening on port 3000")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
