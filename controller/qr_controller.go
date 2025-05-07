// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sat 03/05/25 15.07
// @project api https://github.com/PakaiWA/api/tree/main/controller
//

package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type QRController interface {
	RegisterRoutes(router *httprouter.Router)
	getQRCode(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	showQR(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
