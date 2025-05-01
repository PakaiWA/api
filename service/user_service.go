// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Thu 01/05/25 13.50
// @project api service
//

package service

import (
	"context"
	"github.com/pakaiwa/api/model/api"
)

type UserService interface {
	CreateUser(ctx context.Context, user api.UserRq) api.UserRs
}
