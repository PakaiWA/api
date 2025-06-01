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
		var logFileWriter io.Writer
		var finalWriter io.Writer

		if err := os.MkdirAll(logDir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "[WARN] Fail when create logs dir '%s': %v. Log only shown in console.\n", logDir, err)
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
				fmt.Fprintf(os.Stderr, "[PERINGATAN] Gagal membuka berkas log '%s': %v. Log ke berkas dinonaktifkan.\n", path, err)
				logFileAttemptedAndFailed = true
			} else {
				logFileWriter = outFile
			}
		}

		if logFileWriter != nil {
			finalWriter = io.MultiWriter(os.Stdout, logFileWriter)
		} else {
			finalWriter = os.Stdout
		}

		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		logger = zerolog.New(finalWriter).
			Level(currentLevel).
			With().
			Timestamp().
			Logger()

		logger.Debug().Str("trace_id", config.Get40Space()).Msgf("Inisialisasi logger selesai. Level log: %s.", currentLevel.String())
		if logFileWriter == nil {
			logger.Warn().Str("trace_id", config.Get40Space()).Msg("Logging ke berkas tidak berhasil, log hanya akan tampil di konsol.")
		}
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
