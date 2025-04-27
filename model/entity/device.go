// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sun 27/04/25 17.43
// @project api entity
//

package entity

import "time"

type Device struct {
	Id                 string
	Name               string
	Status             string
	PhoneNumber        string
	CreatedAt          time.Time
	ConnectedAt        time.Time
	DisconnectedAt     time.Time
	DisconnectedReason string
}
