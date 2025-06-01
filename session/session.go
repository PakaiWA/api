// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sat 03/05/25 13.11
// @project api https://github.com/PakaiWA/api/tree/main/session
//

package session

import (
	"context"
	"fmt"
	"github.com/pakaiwa/api/app"
	"os"

	"github.com/mdp/qrterminal/v3"
	"github.com/pakaiwa/pakaiwa"
	"github.com/pakaiwa/pakaiwa/types"
	"github.com/pakaiwa/pakaiwa/types/events"
)

func StartDeviceSession(userJID string) (*pakaiwa.Client, error) {
	jid := types.NewJID(userJID, types.DefaultUserServer)
	container := app.GetContainer()
	fmt.Println("StartDeviceSession Connecting to database...", jid)
	deviceStore := container.NewDevice()
	client := pakaiwa.NewClient(container.NewDevice(), nil)
	fmt.Println("StartDeviceSession Connecting to client...")
	fmt.Println("DeviceStore:", deviceStore.LID, deviceStore.ID)
	fmt.Println("Device Store Client:", client.Store.GetJID(), client.Store.GetLID())

	client.AddEventHandler(eventHandler)

	if client.Store.ID == nil {
		fmt.Println("StartDeviceSession Connecting to Store...")
		qrChan, _ := client.GetQRChannel(context.Background())
		err := client.Connect()
		if err != nil {
			fmt.Println("Error connecting to client:", err)
			return nil, err
		}

		for evt := range qrChan {
			if evt.Event == "code" {
				fmt.Println("Scan this QR:")
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
			} else if evt.Event == "success" || evt.Event == "timeout" {
				break
			}
		}
	} else {
		// auto reconnect
		fmt.Println("StartDeviceSession Reconnecting to client...")
		err := client.Connect()
		if err != nil {
			return nil, err
		}
	}

	return client, nil
}

func eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		fmt.Println("Received a message!", v.Message.GetConversation())
	}
}
