// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sun 04/05/25 21.55
// @project api https://github.com/PakaiWA/api/tree/main/session
//

package session

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/mdp/qrterminal/v3"
	"github.com/pakaiwa/pakaiwa"
)

func QRHandler(client *pakaiwa.Client) string {
	qrChan, _ := client.GetQRChannel(context.Background())

	err := client.Connect()
	if err != nil {
		log.Println("Error connecting client:", err)
		return ""
	}

	qrCodeChan := make(chan string)
	authenticated := make(chan struct{})

	go func() {
		for evt := range qrChan {
			switch evt.Event {
			case "code":
				qrCodeChan <- evt.Code
			case "success":
				log.Println("Client authenticated successfully")
				authenticated <- struct{}{}
				return
			default:
				log.Println("QR event:", evt.Event)
			}
		}
	}()

	qrCode := ""
	select {
	case qrCode = <-qrCodeChan:
		log.Println("QR code received:", qrCode)
		qrterminal.GenerateHalfBlock(qrCode, qrterminal.L, os.Stdout)
	case <-time.After(10 * time.Second):
		log.Println("Failed to receive QR code")
		client.Disconnect()
		return ""
	}

	go func() {
		select {
		case <-authenticated:
			log.Println("user QR code authenticated")
		case <-time.After(30 * time.Second):
			log.Println("QR code not scanned within 30s, disconnecting client...")
			client.Disconnect()
		}
	}()

	return qrCode
}
