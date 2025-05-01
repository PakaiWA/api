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
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
	"github.com/pakaiwa/api/config"
	"github.com/pakaiwa/api/model/api"
	"net/http"
	"strings"
)

type PakaiWAClaim struct {
	Email string `json:"email"`
	Iat   int    `json:"iat"`
	jwt.MapClaims
}

func AuthMiddleware(next httprouter.Handle) httprouter.Handle {
	res := api.ResponseAPI{
		Code:   http.StatusUnauthorized,
		Status: "UNAUTHORIZED",
	}

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			api.WriteToResponseBody(w, res.Code, res)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &PakaiWAClaim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return config.GetJWTKey(), nil
		})

		if err != nil || !token.Valid {
			api.WriteToResponseBody(w, res.Code, res)
			return
		}

		// TODO: check if claims is valid

		// save claims to context
		ctx := context.WithValue(r.Context(), "userEmail", claims.Email)
		next(w, r.WithContext(ctx), ps)
	}
}
