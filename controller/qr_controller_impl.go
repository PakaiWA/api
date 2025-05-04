// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sat 03/05/25 15.08
// @project api controller
//

package controller

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/pakaiwa/api/exception"
	"github.com/pakaiwa/api/helper"
	"github.com/pakaiwa/api/middleware"
	"github.com/pakaiwa/api/model/api"
	"github.com/pakaiwa/api/service"
	"github.com/pakaiwa/api/utils"
	"github.com/skip2/go-qrcode"
	"net/http"
	"net/url"
)

type QRControllerImpl struct {
	QRService service.QRService
}

func NewQRController(QRService service.QRService) QRController {
	return &QRControllerImpl{
		QRService: QRService,
	}
}

func (controller *QRControllerImpl) RegisterRoutes(router *httprouter.Router) {
	router.GET("/qr/:deviceId", middleware.AuthMiddleware(controller.getQRCode))
	router.GET("/qr/show", controller.showQR)
}

func (controller *QRControllerImpl) getQRCode(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Println("Invoke getQRCode Controller")

	qrRs := controller.QRService.GetQRCode(request.Context(), params.ByName("deviceId"))

	apiResponse := api.ResponseAPI{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   qrRs,
		Meta:   nil,
	}

	api.WriteToResponseBody(writer, apiResponse.Code, apiResponse)
}

func (controller *QRControllerImpl) showQR(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Println("Invoke showQR Controller")

	qrCode := request.URL.Query().Get("qrCode")
	if qrCode == "" {
		panic(exception.NewBadRequestError("qrCode query param is required"))
		return
	}

	png, err := qrcode.Encode(qrCode, qrcode.Medium, 256)
	if err != nil {
		helper.PanicIfError(err)
	}

	writer.Header().Set("Content-Type", "image/png")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(png)
	if err != nil {
		helper.PanicIfError(err)
	}
}
