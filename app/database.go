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
	NewLogger().Info().Str("trace_id", config.Get40Space()).Msgf("Connecting to database...")

	onceDb.Do(func() {
		cfg, err := pgxpool.ParseConfig(config.GetDBConn())
		helper.PanicIfError(err)

		cfg.MinConns = config.GetDBMinConn()
		cfg.MaxConns = config.GetDBMaxConn()
		cfg.HealthCheckPeriod = config.GetDBHealthCheckPeriod()

		start := time.Now()
		pool, err = pgxpool.NewWithConfig(ctx, cfg)
		helper.PanicIfError(err)
		NewLogger().Debug().Str("trace_id", config.Get40Space()).Msgf("pgxpool.NewWithConfig took %s", time.Since(start))

		ctx, cancel := context.WithTimeout(ctx, time.Minute)
		defer cancel()
		NewLogger().Info().Str("trace_id", config.Get40Space()).Msg("Pinging database...")
		if err := pool.Ping(ctx); err != nil {
			NewLogger().Error().Str("trace_id", config.Get40Space()).Msgf("Ping timeout: %v", err)
		}
		NewLogger().Info().Str("trace_id", config.Get40Space()).Msg("Pinging done...")
	})

	if pool == nil {
		NewLogger().Error().Str("trace_id", config.Get40Space()).Msgf("Database pool is nil")
	}

	NewLogger().Info().Str("trace_id", config.Get40Space()).Msg("Connected to database...")
	return pool
}
