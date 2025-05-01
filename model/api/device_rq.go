// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sun 27/04/25 17.24
// @project api api
//

package api

type DeviceAddRq struct {
	DeviceId string `validate:"required,min=6,max=100" json:"device_id"`
}
