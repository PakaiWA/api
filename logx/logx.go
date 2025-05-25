// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Min 25/05/25 19.13
// @project api logx
// https://github.com/PakaiWA/api/tree/main/logx
//

package logx

import (
	"github.com/pakaiwa/api/app"
	"github.com/sirupsen/logrus"
)

func Log() *logrus.Logger {
	return app.Log()
}

func Println(args ...interface{}) {
	app.Log().Println(args...)
}

func Debug(args ...interface{}) {
	app.Log().Debug(args...)

}
func Info(args ...interface{}) {
	app.Log().Info(args...)
}

func Warn(args ...interface{}) {
	app.Log().Warn(args...)
}

func Error(args ...interface{}) {
	app.Log().Error(args...)
}

func Fatal(args ...interface{}) {
	app.Log().Fatal(args...)
}

func Printf(format string, args ...interface{}) {
	app.Log().Printf(format, args...)
}

func Infof(format string, args ...interface{}) {
	app.Log().Infof(format, args...)
}
