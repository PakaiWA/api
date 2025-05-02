// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Fri 02/05/25 18.34
// @project api app
//

package app

import (
	"github.com/pakaiwa/api/config"
	"github.com/redis/go-redis/v9"
	"time"
)

var RedisClient *redis.Client

func NewRedisClient() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         config.GetRedisHost(),
		Password:     config.GetRedisPassword(),
		DB:           0,
		DialTimeout:  15 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})
}
