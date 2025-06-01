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
	"github.com/pakaiwa/api/logx"
	"os"
	"time"

	"github.com/mdp/qrterminal/v3"
	"github.com/pakaiwa/pakaiwa"
)

func QRHandler(ctx context.Context, client *pakaiwa.Client) string {
	qrChan, _ := client.GetQRChannel(context.Background())

	err := client.Connect()
	if err != nil {
		logx.Errorf("Error connecting client: %v", err)
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
				logx.Debug("Client authenticated successfully")
				authenticated <- struct{}{}
				return
			default:
				logx.Debugf("QR event: %s", evt.Event)
			}
		}
	}()

	qrCode := ""
	select {
	case qrCode = <-qrCodeChan:
		logx.Debugf("QR code received: %s", qrCode)
		qrterminal.GenerateHalfBlock(qrCode, qrterminal.L, os.Stdout)
	case <-time.After(10 * time.Second):
		logx.Warn("Failed to receive QR code")
		client.Disconnect()
		return ""
	}

	go func() {
		select {
		case <-authenticated:
			logx.InfoCtx(ctx, "user QR code authenticated")
		case <-time.After(30 * time.Second):
			logx.InfoCtx(ctx, "QR code not scanned within 30s, disconnecting client...")
			client.Disconnect()
		}
	}()

	return qrCode
}
