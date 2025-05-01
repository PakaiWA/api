// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Thu 01/05/25 13.48
// @project api controller
//

package controller

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/pakaiwa/api/model/api"
	"github.com/pakaiwa/api/service"
	"net/http"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{UserService: userService}
}

func (controller *UserControllerImpl) RegisterRoutes(router *httprouter.Router) {
	router.POST("/register", controller.CreateUser)
}

func (controller *UserControllerImpl) CreateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Println("Invoke CreateUser Controller")
	req := api.UserRq{}
	api.ReadFromRequestBody(request, &req)

	res := controller.UserService.CreateUser(request.Context(), req)

	apiResponse := api.ResponseAPI{
		Code:   http.StatusCreated,
		Status: "OK",
		Data:   res,
		Meta:   nil,
	}

	api.WriteToResponseBody(writer, apiResponse.Code, apiResponse)
}
