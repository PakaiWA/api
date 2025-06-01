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

func getTraceIDFromContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	traceIDValue := ctx.Value("trace_id")
	if traceIDValue != nil {
		if traceID, ok := traceIDValue.(string); ok {
			return traceID
		}
	}
	return ""
}

func withTraceID(ctx context.Context, event *zerolog.Event) *zerolog.Event {
	traceID := getTraceIDFromContext(ctx)
	if traceID != "" {
		return event.Str("trace_id", traceID)
	}
	return event
}

func Debug(msg string) {
	app.NewLogger().Debug().Msg(msg)
}

func Debugf(format string, args ...interface{}) {
	app.NewLogger().Debug().Msgf(format, args...)
}

func DebugCtx(ctx context.Context, msg string) {
	withTraceID(ctx, app.NewLogger().Debug()).Msg(msg)
}

func DebugfCtx(ctx context.Context, format string, args ...interface{}) {
	withTraceID(ctx, app.NewLogger().Debug()).Msgf(format, args...)
}

func Info(msg string) {
	app.NewLogger().Info().Msg(msg)
}

func Infof(format string, args ...interface{}) {
	app.NewLogger().Info().Msgf(format, args...)
}

func InfoCtx(ctx context.Context, msg string) {
	withTraceID(ctx, app.NewLogger().Info()).Msg(msg)
}

func InfofCtx(ctx context.Context, format string, args ...interface{}) {
	withTraceID(ctx, app.NewLogger().Info()).Msgf(format, args...)
}

func Warn(msg string) {
	app.NewLogger().Warn().Msg(msg)
}

func Warnf(format string, args ...interface{}) {
	app.NewLogger().Warn().Msgf(format, args...)
}

func WarnCtx(ctx context.Context, msg string) {
	withTraceID(ctx, app.NewLogger().Warn()).Msg(msg)
}

func WarnfCtx(ctx context.Context, format string, args ...interface{}) {
	withTraceID(ctx, app.NewLogger().Warn()).Msgf(format, args...)
}

func Error(msg string) {
	app.NewLogger().Error().Msg(msg)
}

func Errorf(format string, args ...interface{}) {
	app.NewLogger().Error().Msgf(format, args...)
}

func ErrorCtx(ctx context.Context, msg string) {
	withTraceID(ctx, app.NewLogger().Error()).Msg(msg)
}

func ErrorfCtx(ctx context.Context, format string, args ...interface{}) {
	withTraceID(ctx, app.NewLogger().Error()).Msgf(format, args...)
}
