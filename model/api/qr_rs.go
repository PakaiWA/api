// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sat 03/05/25 15.40
// @project api https://github.com/PakaiWA/api/tree/main/api
//

package api

type QRCodeRs struct {
	QRCode   string `json:"qr_code"`
	ImageUrl string `json:"image_url"`
}

func SetQrCode(qrCode string) QRCodeRs {
	return QRCodeRs{
		QRCode:   qrCode,
		ImageUrl: "https://api.pakaiwa.com/qr/show/" + qrCode,
	}
}
