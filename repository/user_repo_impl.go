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
	"database/sql"
	"fmt"
	"github.com/pakaiwa/api/helper"
	"github.com/pakaiwa/api/model/entity"
	"github.com/pakaiwa/api/utils"
)

type UserRepoImpl struct{}

func NewUserRepo() UserRepo {
	return &UserRepoImpl{}
}

func (repo UserRepoImpl) CreateUser(ctx context.Context, tx *sql.Tx, user entity.User) entity.User {
	fmt.Println("Invoke CreateUser Repository")

	uuid := utils.GenerateUUID()

	SQL := "insert into management.users (uuid, email, password) values ($1, $2, $3)"
	fmt.Println(SQL, uuid, user.Email, "[REDACTED]")
	_, err := tx.ExecContext(ctx, SQL, uuid, user.Email, user.Password)
	if err != nil {
		helper.PanicIfError(err)
	}

	helper.PanicIfError(err)
	user.Uuid = uuid
	user.Password = ""
	fmt.Println("Success create user", user)
	return user
}

func (repo UserRepoImpl) EmailExist(ctx context.Context, tx *sql.Tx, email string) (bool, error) {
	fmt.Println("Invoke FindByEmail Repository")
	var count int
	SQL := "select count(*) from management.users where email = $1"
	err := tx.QueryRowContext(ctx, SQL, email).Scan(&count)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return true, err
	}

	return count != 0, nil
}
