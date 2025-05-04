// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sun 04/05/25 21.55
// @project api session
//

package session

import (
	"context"
	"github.com/mdp/qrterminal/v3"
	"github.com/pakaiwa/pakaiwa"
	"log"
	"os"
	"time"
)

func QRHandler(client *pakaiwa.Client) string {

	qrChan, _ := client.GetQRChannel(context.Background())
	qrCode := ""

	err := client.Connect()
	if err != nil {
		log.Println("Error connecting client:", err)
		return ""
	}

	timeout := time.After(30 * time.Second)

	for {
		select {
		case evt, ok := <-qrChan:
			if !ok {
				log.Println("QR Channel closed")
				return ""
			}
			if evt.Event == "code" {
				qrCode = evt.Code
				log.Println("QR code received:", qrCode)
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
				return qrCode
			} else {
				log.Println("QR event:", evt.Event)
			}
		case <-timeout:
			log.Println("Timeout waiting for QR code")
			return ""
		}
	}

}
