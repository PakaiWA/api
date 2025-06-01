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
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/pakaiwa/api/config"
	"github.com/rs/zerolog"
)

var (
	once   sync.Once
	logger zerolog.Logger
)

func NewLogger() *zerolog.Logger {
	once.Do(func() {
		logDir := "logs"
		var logFileFailed bool
		var finalWriter io.Writer
		var logFileWriter io.Writer
		currentLevel := getLogLevel()

		// Logging to file
		if err := os.MkdirAll(logDir, 0755); err != nil {
			logger.Warn().Str("trace_id", config.Get40Space()).Msgf("[WARN] Failed when creating log dir %s:%v. The log is only displayed in the console.", logDir, err)
			logFileFailed = true
		} else {
			date := time.Now().Format("20060102")
			idx := 0
			var path string
			for {
				path = filepath.Join(logDir, fmt.Sprintf("%s-%02d.log", date, idx))
				info, err := os.Stat(path)
				if os.IsNotExist(err) || (err == nil && info.Size() < 1<<30) { // 1GB limit
					break
				}
				idx++
			}

			outFile, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			if err != nil {
				logger.Warn().Str("trace_id", config.Get40Space()).Msgf("[WARN] Failed to open log file %s: %v. Logging to the file is disabled.", path, err)
				logFileFailed = true
			} else {
				logFileWriter = outFile
			}
		}

		// Console logging
		consoleWriter := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}

		consoleWriter.PartsOrder = []string{
			zerolog.TimestampFieldName,
			zerolog.LevelFieldName,
			"trace_id",
			zerolog.MessageFieldName,
		}

		consoleWriter.FieldsExclude = []string{"trace_id"}

		if logFileFailed {
			logger.Warn().Str("trace_id", config.Get40Space()).Msg("Logging to the file is not successful, the log will only appear on the console.")
			finalWriter = consoleWriter
		} else {
			fmt.Println("INIT ALL")
			finalWriter = io.MultiWriter(consoleWriter, logFileWriter)
		}

		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		logger = zerolog.New(finalWriter).
			Level(currentLevel).
			With().
			Timestamp().
			Logger()

		logger.Debug().Str("trace_id", config.Get40Space()).Msgf("Logger initialization complete. Log level: %s.", currentLevel.String())
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
