// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Wed 30/04/25 23.06
// @project api https://github.com/PakaiWA/api/tree/main/middleware
//

package middleware

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
	"github.com/pakaiwa/api/app"
	"github.com/pakaiwa/api/config"
	"github.com/pakaiwa/api/logx"
	"github.com/pakaiwa/api/model/api"
	"github.com/pakaiwa/api/utils"
	"github.com/redis/go-redis/v9"
	"net/http"
	"strconv"
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
		traceID := r.Header.Get("ax-request-id")
		if traceID == "" {
			traceID = utils.GenerateUUID()
		}

		ctx := context.WithValue(r.Context(), "trace_id", traceID)

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

		exp, err := app.RedisClient.Get(context.Background(), claims.Email).Result()
		if errors.Is(err, redis.Nil) {
			logx.Info(ctx, "Token is valid")
		} else if err != nil {
			logx.Infof("Error when retrieve data from Redis: %v", err)
			res.Code = http.StatusInternalServerError
			res.Status = "INTERNAL_SERVER_ERROR"
			api.WriteToResponseBody(w, res.Code, res)
			return
		} else {
			if logoutOn, _ := strconv.Atoi(exp); logoutOn > claims.Iat {
				logx.Info(ctx, "Token is invalid")
				res.Code = http.StatusUnauthorized
				res.Status = "UNAUTHORIZED"
				api.WriteToResponseBody(w, res.Code, res)
				return
			}
		}

		// save claims to context
		ctx = context.WithValue(ctx, "userEmail", claims.Email)
		next(w, r.WithContext(ctx), ps)
	}
}

func AdminMiddleware(next httprouter.Handle) httprouter.Handle {
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

		if tokenString != config.GetAdminToken() {
			api.WriteToResponseBody(w, res.Code, res)
			return
		}

		next(w, r, ps)
	}
}
