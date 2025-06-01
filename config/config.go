// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sun 27/04/25 15.43
// @project api https://github.com/PakaiWA/api/tree/main/config
//

package config

import (
	"github.com/spf13/viper"
	"time"
)

func GetDBConn() string {
	return viper.GetString("db.pakaiwa.host")
}

func GetDBMinConn() int32 {
	minConn := viper.GetInt32("db.pakaiwa.MinConns")
	if minConn <= 0 {
		minConn = 1
	}
	return minConn
}

func GetDBMaxConn() int32 {
	maxConn := viper.GetInt32("db.pakaiwa.MaxConns")
	if maxConn <= 0 {
		maxConn = 10
	}
	return maxConn
}

func GetDBHealthCheckPeriod() time.Duration {
	if viper.IsSet("db.pakaiwa.HealthCheckPeriod") {
		val := viper.GetDuration("db.pakaiwa.HealthCheckPeriod")
		if val > 0 {
			return val
		}
	}
	return 1 * time.Minute
}

func GetJWTKey() []byte { return []byte(viper.GetString("app.jwt.sign_key")) }

func GetAdminToken() string { return viper.GetString("app.admin.token") }

func GetAllDevicesSQL() string { return viper.GetString("app.sql.getAllDevices") }

func GetDeviceByIdSQL() string { return viper.GetString("app.sql.getDeviceById") }

func GetDeleteDeviceSQL() string { return viper.GetString("app.sql.deleteDeviceById") }

func GetAddDeviceSQL() string { return viper.GetString("app.sql.addDevice") }

func GetCountDeviceSQL() string { return viper.GetString("app.sql.countDeviceById") }

func GetRedisHost() string { return viper.GetString("redis.host") }

func GetRedisPassword() string { return viper.GetString("redis.password") }

func GetLogLevel() string { return viper.GetString("log.level") }
