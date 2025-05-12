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
	pakaiWAMap = make(map[string]*pakaiwa.Client)
	pwaLock    = sync.Mutex{}
	mutex      = sync.RWMutex{}
)

func NewDevicePakaiWA(deviceId string) *pakaiwa.Client {
	fmt.Println("Invoke NewDevicePakaiWA session")
	pwaLock.Lock()
	defer pwaLock.Unlock()

	if existingClient, ok := pakaiWAMap[deviceId]; ok {
		if existingClient.IsConnected() {
			log.Println("Client already connected for device:", deviceId)
		}
	}

	container := app.GetContainer()
	store := container.NewDevice()
	client := pakaiwa.NewClient(store, waLog.Stdout("PakaiWA", "INFO", true))

	pakaiWAMap[deviceId] = client
	return client
}

func RegisterClient(uuid string, client *pakaiwa.Client) {
	mutex.Lock()
	defer mutex.Unlock()
	fmt.Println("Register client:", uuid)
	pakaiWAMap[uuid] = client
}

func GetClient(uuid string) (*pakaiwa.Client, bool) {
	mutex.RLock()
	defer mutex.RUnlock()
	client, exists := pakaiWAMap[uuid]
	return client, exists
}

func RemoveClient(uuid string) {
	mutex.Lock()
	defer mutex.Unlock()
	delete(pakaiWAMap, uuid)
}

func ListClients() map[string]*pakaiwa.Client {
	mutex.RLock()
	defer mutex.RUnlock()
	clone := make(map[string]*pakaiwa.Client)
	for k, v := range pakaiWAMap {
		clone[k] = v
	}
	return clone
}
