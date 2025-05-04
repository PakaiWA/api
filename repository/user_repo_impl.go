// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sun 27/04/25 18.06
// @project api repository
//

package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/pakaiwa/api/helper"
	"github.com/pakaiwa/api/model/entity"
	"github.com/pakaiwa/api/utils"
	"strings"
)

type UserRepoImpl struct{}

func NewUserRepo() UserRepo {
	return &UserRepoImpl{}
}

func (repo UserRepoImpl) CreateUser(ctx context.Context, tx pgx.Tx, user entity.User) entity.User {
	fmt.Println("Invoke CreateUser Repository")

	uuid := utils.GenerateUUID()
	userEmail := strings.ToLower(user.Email)

	SQL := "insert into management.users (uuid, email, password) values ($1, $2, $3)"
	fmt.Println(SQL, uuid, userEmail, "[REDACTED]")
	_, err := tx.Exec(ctx, SQL, uuid, userEmail, user.Password)
	if err != nil {
		helper.PanicIfError(err)
	}

	helper.PanicIfError(err)
	user.Uuid = uuid
	user.Password = "" // Clear sensitive data
	fmt.Printf("Success create user with UUID: %s and Email: %s\n", user.Uuid, userEmail)
	return user
}

func (repo UserRepoImpl) EmailExist(ctx context.Context, tx pgx.Tx, email string) (bool, error) {
	fmt.Println("Invoke FindByEmail Repository")
	var count int
	userEmail := strings.ToLower(email)

	SQL := "select count(*) from management.users where email = $1"
	err := tx.QueryRow(ctx, SQL, userEmail).Scan(&count)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return true, err
	}

	return count != 0, nil
}

func (repo UserRepoImpl) Login(ctx context.Context, tx pgx.Tx, email, pass string) (entity.User, error) {
	fmt.Println("Invoke Login Repository")
	SQL := "SELECT uuid, email, password FROM management.users WHERE email = $1"

	user := entity.User{}
	userEmail := strings.ToLower(email)

	err := tx.QueryRow(ctx, SQL, userEmail).Scan(&user.Uuid, &user.Email, &user.Password)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return user, err
	}

	if !utils.CheckPasswordHash(pass, user.Password) {
		fmt.Println("Invalid password")
		return user, fmt.Errorf("invalid password")
	}

	user.Password = "" // Clear sensitive data
	fmt.Printf("Success login user with UUID: %s and Email: %s\n", user.Uuid, userEmail)
	return user, nil
}
