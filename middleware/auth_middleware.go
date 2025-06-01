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
	"github.com/pakaiwa/api/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
	"github.com/pakaiwa/api/app"
	"github.com/pakaiwa/api/config"
	"github.com/pakaiwa/api/logx"
	"github.com/pakaiwa/api/model/api"
	"github.com/redis/go-redis/v9"
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
		requestCtx := context.WithValue(r.Context(), "trace_id", getTraceID(r))

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

		exp, errRedis := app.RedisClient.Get(requestCtx, claims.Email).Result()
		if errors.Is(errRedis, redis.Nil) {
			logx.InfoCtx(requestCtx, "Token is valid")
		} else if errRedis != nil {
			logx.InfofCtx(requestCtx, "Error when retrieve data from Redis: %v", errRedis)
			res.Code = http.StatusInternalServerError
			res.Status = "INTERNAL_SERVER_ERROR"
			api.WriteToResponseBody(w, res.Code, res)
			return
		} else {
			if logoutOn, _ := strconv.Atoi(exp); logoutOn > claims.Iat {
				logx.InfoCtx(requestCtx, "Token is invalid")
				res.Code = http.StatusUnauthorized
				res.Status = "UNAUTHORIZED"
				api.WriteToResponseBody(w, res.Code, res)
				return
			}
		}

		ctxWithUserEmail := context.WithValue(requestCtx, "userEmail", claims.Email)
		next(w, r.WithContext(ctxWithUserEmail), ps)
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
		ctx := context.WithValue(r.Context(), "trace_id", getTraceID(r))
		next(w, r.WithContext(ctx), ps)
	}
}

func getTraceID(r *http.Request) string {
	traceID := r.Header.Get("ax-request-id")
	if traceID == "" {
		traceID = utils.GenerateUUID()
	}
	return traceID
}
