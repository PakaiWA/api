// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sun 27/04/25 15.43
// @project api config
//

package config

import (
	"github.com/spf13/viper"
)

func GetDBCon() string {
	return viper.GetString("db.pakaiwa.host")
}

func GetJWTKey() []byte {
	return []byte(viper.GetString("app.jwt.sign_key"))
}

func GetAdminToken() string {
	return viper.GetString("app.admin.token")
}

func GetAllDevicesSQL() string {
	return viper.GetString("app.sql.getAllDevices")
}

func GetDeviceByIdSQL() string {
	return viper.GetString("app.sql.getDeviceById")
}

func GetDeleteDeviceSQL() string {
	return viper.GetString("app.sql.deleteDeviceById")
}

func GetAddDeviceSQL() string {
	return viper.GetString("app.sql.addDevice")
}

func GetCountDeviceSQL() string {
	return viper.GetString("app.sql.countDeviceById")
}

func GetRedisHost() string {
	return viper.GetString("redis.host")
}

func GetRedisPassword() string {
	return viper.GetString("redis.password")
}
