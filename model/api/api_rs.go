// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sun 27/04/25 17.22
// @project api https://github.com/PakaiWA/api/tree/main/api
//

package api

import (
	"encoding/json"
	"net/http"

	"github.com/pakaiwa/api/helper"
)

type ResponseAPI struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Meta   *Meta       `json:"meta,omitempty"`
}

type Meta struct {
	LastKey  string `json:"last_key,omitempty"`
	Location string `json:"location,omitempty"`
}

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	helper.PanicIfError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, statusCode int, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	helper.PanicIfError(err)
}
