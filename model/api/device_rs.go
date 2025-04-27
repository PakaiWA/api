// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sun 27/04/25 17.27
// @project api api
//

package api

import "time"

type DeviceRs struct {
	Id                 string    `json:"id"`
	Status             string    `json:"status"`
	PhoneNumber        string    `json:"phone_number"`
	CreatedAt          time.Time `json:"created_at"`
	ConnectedAt        time.Time `json:"connected_at"`
	DisconnectedAt     time.Time `json:"disconnected_at"`
	DisconnectedReason string    `json:"disconnected_reason"`
}
