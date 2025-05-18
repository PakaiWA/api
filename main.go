// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sun 27/04/25 17.10
// @project api https://github.com/PakaiWA/api/tree/main/main
//

package main

import (
	"context"
	"github.com/KAnggara75/scc2go"
	"github.com/go-playground/validator/v10"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/julienschmidt/httprouter"
	"github.com/pakaiwa/api/app"
	"github.com/pakaiwa/api/controller"
	"github.com/pakaiwa/api/exception"
	"github.com/pakaiwa/api/helper"
	logger "github.com/pakaiwa/api/log"
	"github.com/pakaiwa/api/repository"
	"github.com/pakaiwa/api/service"
	"github.com/pakaiwa/api/usecase"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

var log *logrus.Logger

func init() {
	scc2go.GetEnv(os.Getenv("SCC_URL"), os.Getenv("AUTH"), true)
	app.NewRedisClient()
	log = logger.Logger()
	//session.RestoreAllClient()
}

func main() {
	log.Debug("Starting server")
	db := app.NewDBConn(context.Background())
	router := httprouter.New()
	log.Debug("Initializing router...")
	validate := validator.New()
	log.Debug("Initializing validator...")

	registerControllers(router, db, validate)

	server := &http.Server{
		Addr:         "localhost:3000",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	router.PanicHandler = exception.ErrorHandler
	log.Infof("Listening on => https://%s", server.Addr)
	log.Infof("PakaiWA running at %s", time.Now().Format(time.RFC850))
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
