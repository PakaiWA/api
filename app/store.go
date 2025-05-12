// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sun 11/05/25 19.06
// @project api app
// https://github.com/PakaiWA/api/tree/main/app
//

package app

import (
	"github.com/pakaiwa/api/config"
	"github.com/pakaiwa/pakaiwa/store/sqlstore"
	waLog "github.com/pakaiwa/pakaiwa/util/log"
	"log"
	"sync"
)

var (
	container *sqlstore.Container
	onceStore sync.Once
)

// GetContainer returns a singleton instance of sqlstore.Container
func GetContainer() *sqlstore.Container {
	onceStore.Do(func() {
		dbLog := waLog.Stdout("Database", "DEBUG", true)
		var err error
		container, err = sqlstore.New(config.GetDBCon(), dbLog)
		if err != nil {
			log.Fatalf("Failed to initialize sqlstore container: %v", err)
		}
	})
	return container
}
