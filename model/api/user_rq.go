// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Thu 01/05/25 13.54
// @project api https://github.com/PakaiWA/api/tree/main/api
//

package api

type UserRq struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=6,max=100" json:"password"`
}
