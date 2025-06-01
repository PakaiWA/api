// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sun 18/05/25 18.03
// @project api log
// https://github.com/PakaiWA/api/tree/main/log
//

package app

import (
	"fmt"
	"github.com/pakaiwa/api/config"
	"github.com/rs/zerolog"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	once   sync.Once
	logger zerolog.Logger
)

func NewLogger() *zerolog.Logger {
	once.Do(func() {
		logDir := "logs"
		_ = os.MkdirAll(logDir, os.ModePerm)

		date := time.Now().Format("20060102")
		idx := 0
		var path string
		for {
			path = filepath.Join(logDir, fmt.Sprintf("%s-%02d.log", date, idx))
			info, err := os.Stat(path)
			if os.IsNotExist(err) || (err == nil && info.Size() < 1<<30) {
				break
			}
			idx++
		}

		outFile, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			panic(fmt.Sprintf("failed to open log file: %v", err))
		}

		writer := io.MultiWriter(os.Stdout, outFile)
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		level := getLogLevel()
		logger = zerolog.New(writer).
			Level(level).
			With().
			Timestamp().
			Logger()

		logger.Info().Msgf("Log level set to: %s", level)
	})

	return &logger
}

func getLogLevel() zerolog.Level {
	level, err := zerolog.ParseLevel(strings.ToLower(config.GetLogLevel()))
	if err != nil {
		level = zerolog.InfoLevel
	}
	return level
}
