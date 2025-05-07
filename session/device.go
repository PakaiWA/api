// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sun 04/05/25 22.00
// @project api https://github.com/PakaiWA/api/tree/main/session
//

package session

import (
	"fmt"
	"log"
	"sync"

	"github.com/pakaiwa/api/config"
	"github.com/pakaiwa/api/helper"
	"github.com/pakaiwa/pakaiwa"
	"github.com/pakaiwa/pakaiwa/store/sqlstore"
	waLog "github.com/pakaiwa/pakaiwa/util/log"
)

var (
	clientMap  = make(map[string]*pakaiwa.Client)
	clientLock = sync.Mutex{}
)

func NewDevicePakaiWA(deviceId string) *pakaiwa.Client {
	fmt.Println("Invoke NewDevicePakaiWA session")
	clientLock.Lock()
	defer clientLock.Unlock()

	if existingClient, ok := clientMap[deviceId]; ok {
		if existingClient.IsConnected() {
			log.Println("Client already connected for device:", deviceId)
		}
	}

	// Load device and store
	dbLog := waLog.Stdout("DB", "DEBUG", true)
	container, err := sqlstore.New(config.GetDBCon(), dbLog)
	if err != nil {
		helper.PanicIfError(err)
	}

	store := container.NewDevice()
	client := pakaiwa.NewClient(store, waLog.Stdout("PakaiWA", "DEBUG", true))

	clientMap[deviceId] = client

	return client
}
