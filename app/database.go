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
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pakaiwa/api/config"
	"github.com/pakaiwa/api/helper"
	"sync"
)

var (
	pool *pgxpool.Pool
	once sync.Once
)

func NewDBConn(ctx context.Context) *pgxpool.Pool {
	once.Do(func() {
		var err error
		pool, err = pgxpool.New(ctx, config.GetDBCon())
		helper.PanicIfError(err)
	})

	// Verify the connection
	if err := pool.Ping(ctx); err != nil {
		helper.PanicIfError(err)
	}

	return pool
}
