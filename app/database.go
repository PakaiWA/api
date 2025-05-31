// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sun 27/04/25 21.28
// @project api https://github.com/PakaiWA/api/tree/main/app
//

package app

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"sync"
	"time"

	"github.com/pakaiwa/api/config"
	"github.com/pakaiwa/api/helper"
)

var (
	pool   *pgxpool.Pool
	onceDb sync.Once
)

func NewDBConn(ctx context.Context) *pgxpool.Pool {
	NewLogger().Info().Msgf("Connecting to database...")

	onceDb.Do(func() {
		cfg, err := pgxpool.ParseConfig(config.GetDBCon())
		helper.PanicIfError(err)

		cfg.MinConns = 1
		cfg.MaxConns = 10
		cfg.HealthCheckPeriod = time.Minute

		start := time.Now()
		pool, err = pgxpool.NewWithConfig(ctx, cfg)
		helper.PanicIfError(err)
		NewLogger().Debug().Msgf("pgxpool.NewWithConfig took %s", time.Since(start))

		//ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		//defer cancel()
		NewLogger().Info().Msg("Pinging database...")
		if err := pool.Ping(ctx); err != nil {
			NewLogger().Error().Msgf("Ping timeout: %v", err)
		}
		NewLogger().Info().Msg("Pinging done...")
	})

	if pool == nil {
		NewLogger().Error().Msgf("Database pool is nil")
	}

	NewLogger().Info().Msg("Connected to database...")
	return pool
}
