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
	"github.com/pakaiwa/api/app"
	"log"
	"sync"

	"github.com/pakaiwa/pakaiwa"
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

	container := app.GetContainer()
	store := container.NewDevice()
	client := pakaiwa.NewClient(store, waLog.Stdout("PakaiWA", "INFO", true))

	clientMap[deviceId] = client

	return client
}
