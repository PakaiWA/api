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
	"github.com/rs/zerolog"
)

func Debug(msg string, fields ...func(e *zerolog.Event) *zerolog.Event) {
	e := app.NewLogger().Debug()
	for _, f := range fields {
		e = f(e)
	}
	e.Msg(msg)
}

func Debugf(format string, args ...interface{}) {
	app.NewLogger().Debug().Msgf(format, args...)
}

func Info(msg string, fields ...func(e *zerolog.Event) *zerolog.Event) {
	e := app.NewLogger().Info()
	for _, f := range fields {
		e = f(e)
	}
	e.Msg(msg)
}

func Infof(format string, args ...interface{}) {
	app.NewLogger().Info().Msgf(format, args...)
}

func Warn(msg string, fields ...func(e *zerolog.Event) *zerolog.Event) {
	e := app.NewLogger().Warn()
	for _, f := range fields {
		e = f(e)
	}
	e.Msg(msg)
}

func Error(msg string, fields ...func(e *zerolog.Event) *zerolog.Event) {
	e := app.NewLogger().Error()
	for _, f := range fields {
		e = f(e)
	}
	e.Msg(msg)
}

func Errorf(format string, args ...interface{}) {
	app.NewLogger().Error().Msgf(format, args...)
}

func Fatal(msg string, fields ...func(e *zerolog.Event) *zerolog.Event) {
	e := app.NewLogger().Fatal()
	for _, f := range fields {
		e = f(e)
	}
	e.Msg(msg)
}
