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
	"fmt"
	"github.com/spf13/viper"
)

func GetDBCon() string {
	fmt.Println(viper.Get("db.pakaiwa.host"))
	return viper.GetString("db.pakaiwa.host")
}
