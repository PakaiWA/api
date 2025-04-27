// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sun 27/04/25 21.28
// @project api app
//

package app

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pakaiwa/api/config"
	"github.com/pakaiwa/api/helper"
	"time"
)

func NewDBConn() *sql.DB {
	db, err := sql.Open("pgx", config.GetDBCon())
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
