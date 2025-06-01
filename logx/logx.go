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
	"context"
	"github.com/pakaiwa/api/app"
	"github.com/rs/zerolog"
)

var logger = app.NewLogger()

func Debug(msg string, fields ...func(e *zerolog.Event) *zerolog.Event) {
	e := logger.Debug()
	for _, f := range fields {
		e = f(e)
	}
	e.Msg(msg)
}

func Debugf(format string, args ...interface{}) {
	logger.Debug().Msgf(format, args...)
}

func Info(ctx context.Context, msg string, fields ...func(e *zerolog.Event) *zerolog.Event) {
	e := logger.Info().Str("trace_id", ctx.Value("trace_id").(string))
	for _, f := range fields {
		e = f(e)
	}
	e.Msg(msg)
}

func Infof(format string, args ...interface{}) {
	logger.Info().Msgf(format, args...)
}

func Warn(msg string, fields ...func(e *zerolog.Event) *zerolog.Event) {
	e := logger.Warn()
	for _, f := range fields {
		e = f(e)
	}
	e.Msg(msg)
}

func Error(msg string, fields ...func(e *zerolog.Event) *zerolog.Event) {
	e := logger.Error()
	for _, f := range fields {
		e = f(e)
	}
	e.Msg(msg)
}

func Errorf(format string, args ...interface{}) {
	logger.Error().Msgf(format, args...)
}

func Fatal(msg string, fields ...func(e *zerolog.Event) *zerolog.Event) {
	e := logger.Fatal()
	for _, f := range fields {
		e = f(e)
	}
	e.Msg(msg)
}
