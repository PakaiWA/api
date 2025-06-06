// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sun 04/05/25 22.04
// @project api https://github.com/PakaiWA/api/tree/main/usecase
//

package usecase

import (
	"context"
	"net/http"

	"github.com/pakaiwa/api/model/api"
)

type QRUsecase interface {
	GetQRCode(ctx context.Context, request *http.Request) api.QRCodeRs
}
