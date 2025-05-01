// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Thu 01/05/25 13.58
// @project api service
//

package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/pakaiwa/api/exception"
	"github.com/pakaiwa/api/helper"
	"github.com/pakaiwa/api/model/api"
	"github.com/pakaiwa/api/model/entity"
	"github.com/pakaiwa/api/repository"
	"github.com/pakaiwa/api/utils"
)

type UserServiceImpl struct {
	Repo     repository.UserRepo
	DB       *sql.DB
	Validate *validator.Validate
}

func NewUserService(repo repository.UserRepo, db *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		Repo:     repo,
		DB:       db,
		Validate: validate,
	}
}

func (service UserServiceImpl) Login(ctx context.Context, req api.UserRq) api.UserRs {
	fmt.Println("Invoke Login Service")
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.Repo.Login(ctx, tx, req.Email, req.Password)
	if err != nil {
		fmt.Println("Error logging in:", err)
		helper.PanicIfError(err)
	}

	return api.ToUserResponse(user)
}

func (service UserServiceImpl) CreateUser(ctx context.Context, req api.UserRq) api.UserRs {
	fmt.Println("Invoke CreateUser Service")
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	exist, err := service.Repo.EmailExist(ctx, tx, req.Email)
	if err != nil {
		fmt.Println("Error checking email existence:", err)
		helper.PanicIfError(err)
	}

	if exist {
		fmt.Println("Email already exists")
		panic(exception.NewBadRequestError("Email already registered"))
	}

	pass, _ := utils.HashPassword(req.Password)

	user := entity.User{
		Uuid:     utils.GenerateUUID(),
		Email:    req.Email,
		Password: pass,
	}

	user = service.Repo.CreateUser(ctx, tx, user)

	return api.ToUserResponse(user)

}
