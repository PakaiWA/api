// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Wed 30/04/25 23.06
// @project api middleware
//

package middleware

import (
	"github.com/julienschmidt/httprouter"
	"github.com/pakaiwa/api/model/api"
	"net/http"
	"strings"
)

func AuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		token := r.Header.Get("Authorization")
		token = strings.TrimPrefix(token, "Bearer ")
		if token != "RAHASIA" {
			res := api.ResponseAPI{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
			}
			api.WriteToResponseBody(w, res.Code, res)
			return
		}
		next(w, r, ps)
	}
}
