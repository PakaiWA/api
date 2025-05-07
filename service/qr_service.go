// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sat 03/05/25 15.51
// @project api https://github.com/PakaiWA/api/tree/main/service
//

package service

import (
	"context"

	"github.com/pakaiwa/api/model/api"
)

type QRService interface {
	GetQRCode(ctx context.Context, deviceId string) api.QRCodeRs
}
