// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Thu 01/05/25 13.58
// @project api https://github.com/PakaiWA/api/tree/main/service
//

package service

import (
	"context"
	"github.com/pakaiwa/api/logx"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pakaiwa/api/app"
	"github.com/pakaiwa/api/exception"
	"github.com/pakaiwa/api/helper"
	"github.com/pakaiwa/api/model/api"
	"github.com/pakaiwa/api/model/entity"
	"github.com/pakaiwa/api/repository"
	"github.com/pakaiwa/api/utils"
)

type UserServiceImpl struct {
	Repo     repository.UserRepo
	DB       *pgxpool.Pool
	Validate *validator.Validate
}

func NewUserService(repo repository.UserRepo, db *pgxpool.Pool, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		Repo:     repo,
		DB:       db,
		Validate: validate,
	}
}

func (service UserServiceImpl) Logout(ctx context.Context) {
	logx.DebugCtx(ctx, "Invoke Logout Service")
	tx, conn, err := helper.DBTransaction(ctx, service.DB)
	logx.DebugCtx(ctx, "Invoke DBTransaction")
	helper.PanicIfError(err)
	defer conn.Release()
	defer helper.CommitOrRollback(ctx, tx)

	email := ctx.Value("userEmail").(string)

	exist, err := service.Repo.EmailExist(ctx, tx, email)
	if err != nil {
		logx.ErrorfCtx(ctx, "Error checking email existence:", err)
		helper.PanicIfError(err)
	}

	if !exist {
		panic(exception.NewHTTPError(http.StatusBadRequest, "invalid token: email not registered"))
	}

	app.RedisClient.Set(ctx, email, time.Now().Unix(), 0)
}

func (service UserServiceImpl) Login(ctx context.Context, req api.UserRq) api.UserRs {
	logx.DebugCtx(ctx, "Invoke Login Service")
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)

	tx, conn, err := helper.DBTransaction(ctx, service.DB)
	logx.DebugCtx(ctx, "Invoke DBTransaction")
	helper.PanicIfError(err)
	defer conn.Release()
	defer helper.CommitOrRollback(ctx, tx)

	user, err := service.Repo.Login(ctx, tx, req.Email, req.Password)
	if err != nil {
		logx.ErrorfCtx(ctx, "Error logging in:", err)
		panic(exception.NewHTTPError(http.StatusUnauthorized, "Wrong email or password"))
	}

	return api.ToUserResponse(ctx, user)
}

func (service UserServiceImpl) CreateUser(ctx context.Context, req api.UserRq) api.UserRs {
	logx.DebugCtx(ctx, "Invoke CreateUser Service")
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)

	tx, conn, err := helper.DBTransaction(ctx, service.DB)
	logx.DebugCtx(ctx, "Invoke DBTransaction")
	helper.PanicIfError(err)
	defer conn.Release()
	defer helper.CommitOrRollback(ctx, tx)

	exist, err := service.Repo.EmailExist(ctx, tx, req.Email)
	if err != nil {
		logx.ErrorfCtx(ctx, "Error checking email existence:", err)
		helper.PanicIfError(err)
	}

	if exist {
		panic(exception.NewHTTPError(http.StatusBadRequest, "Email already registered"))
	}

	pass, _ := utils.HashPassword(req.Password)

	user := entity.User{
		Uuid:     utils.GenerateUUID(),
		Email:    req.Email,
		Password: pass,
	}

	user = service.Repo.CreateUser(ctx, tx, user)

	return api.ToUserResponse(ctx, user)

}
