// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Thu 01/05/25 16.28
// @project api https://github.com/PakaiWA/api/tree/main/repository
//

package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/pakaiwa/api/model/entity"
)

type UserRepo interface {
	CreateUser(ctx context.Context, tx pgx.Tx, user entity.User) entity.User
	EmailExist(ctx context.Context, tx pgx.Tx, email string) (bool, error)
	Login(ctx context.Context, tx pgx.Tx, email, pass string) (entity.User, error)
}
