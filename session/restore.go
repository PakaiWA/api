// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sun 11/05/25 18.50
// @project api session
// https://github.com/PakaiWA/api/tree/main/session
//

package session

import (
	"github.com/pakaiwa/api/config"
	"github.com/pakaiwa/pakaiwa"
	"github.com/pakaiwa/pakaiwa/store"
	"github.com/pakaiwa/pakaiwa/store/sqlstore"
	waLog "github.com/pakaiwa/pakaiwa/util/log"
	"log"
	"time"
)

func RestoreAllClient() {
	dbLog := waLog.Stdout("Database", "INFO", true)
	container, err := sqlstore.New(config.GetDBCon(), dbLog)
	if err != nil {
		panic(err)
	}

	deviceStore, err := container.GetAllDevices()
	if err != nil {
		panic(err)
	}

	clientLog := waLog.Stdout("Client", "INFO", true)

	var successCount, failCount int

	for _, device := range deviceStore {
		go func(device *store.Device) {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("Recovered while restoring device %s: %v", device.ID, r)
					failCount++
				}
			}()

			client := pakaiwa.NewClient(device, clientLog)
			client.AddEventHandler(eventHandler)
			if err := client.Connect(); err != nil {
				log.Printf("Error connecting device %s: %v", device.ID, err)
				failCount++
				return
			}
			successCount++
		}(device)
	}

	time.Sleep(5 * time.Second)

	log.Printf("Restore complete: %d success, %d failed", successCount, failCount)
}
