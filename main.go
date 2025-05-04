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
	"context"
	"fmt"
	"github.com/KAnggara75/scc2go"
	"github.com/go-playground/validator/v10"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/julienschmidt/httprouter"
	"github.com/pakaiwa/api/app"
	"github.com/pakaiwa/api/controller"
	"github.com/pakaiwa/api/exception"
	"github.com/pakaiwa/api/helper"
	"github.com/pakaiwa/api/repository"
	"github.com/pakaiwa/api/service"
	"github.com/pakaiwa/api/usecase"
	"net/http"
	"os"
)

var ctx = context.Background()

func init() {
	scc2go.GetEnv(os.Getenv("SCC_URL"), os.Getenv("AUTH"))
	app.NewRedisClient()
}

func main() {
	db := app.NewDBConn(ctx)
	router := httprouter.New()
	validate := validator.New()

	registerControllers(router, db, validate)
	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	fmt.Printf("Listening on => https://%s", server.Addr)
	helper.PanicIfError(server.ListenAndServe())
}

func registerControllers(router *httprouter.Router, db *pgxpool.Pool, validate *validator.Validate) {
	// Device
	deviceRepo := repository.NewDeviceRepository()
	deviceService := service.NewDeviceService(deviceRepo, db, validate)
	deviceController := controller.NewDeviceController(deviceService)
	deviceController.RegisterRoutes(router)

	// User
	userRepo := repository.NewUserRepo()
	userService := service.NewUserService(userRepo, db, validate)
	userController := controller.NewUserController(userService)
	userController.RegisterRoutes(router)

	// QR
	qrDeviceRepo := repository.NewDeviceRepository() // kalau bisa reuse deviceRepo
	qrService := service.NewQRService(qrDeviceRepo, db)
	qrUsecase := usecase.NewQRUsecase(qrService, deviceService)
	qrController := controller.NewQRController(qrUsecase)
	qrController.RegisterRoutes(router)
}
