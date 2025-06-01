// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sun 27/04/25 17.19
// @project api https://github.com/PakaiWA/api/tree/main/exception
//

package exception

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/pakaiwa/api/model/api"
)

type HTTPError struct {
	Code   int
	Status string
	Msg    string
}

func NewHTTPError(code int, msg string) HTTPError {
	return HTTPError{Code: code, Msg: msg}
}

func (e HTTPError) Error() string {
	return e.Msg
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	apiRs := api.ResponseAPI{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   fmt.Sprintf("%v", err),
	}

	switch e := err.(type) {
	case validator.ValidationErrors:
		apiRs.Code = http.StatusBadRequest
		apiRs.Status = "BAD REQUEST"
		apiRs.Data = e.Error()
	case HTTPError:
		apiRs.Data = e.Msg
		apiRs.Code = e.Code
		apiRs.Status = http.StatusText(e.Code)
	default:
		apiRs.Code = http.StatusInternalServerError
	}

	api.WriteToResponseBody(w, apiRs.Code, apiRs)
}
