// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sun 27/04/25 17.27
// @project api https://github.com/PakaiWA/api/tree/main/api
//

package api

import (
	"github.com/pakaiwa/api/logx"
	"time"

	"github.com/pakaiwa/api/model/entity"
)

type DeviceRs struct {
	Id                 string     `json:"id"`
	Status             string     `json:"status"`
	PhoneNumber        string     `json:"phone_number,omitempty"`
	CreatedAt          time.Time  `json:"created_at"`
	ConnectedAt        *time.Time `json:"connected_at,omitempty"`
	DisconnectedAt     *time.Time `json:"disconnected_at,omitempty"`
	DisconnectedReason *string    `json:"disconnected_reason,omitempty"`
}

func ToDeviceResponse(device entity.Device) DeviceRs {
	logx.Debug("Invoke ToDeviceResponse")
	return DeviceRs{
		Id:                 device.Name,
		Status:             device.Status,
		PhoneNumber:        device.PhoneNumber,
		CreatedAt:          device.CreatedAt,
		ConnectedAt:        device.ConnectedAt,
		DisconnectedAt:     device.DisconnectedAt,
		DisconnectedReason: device.DisconnectedReason,
	}
}

func ToDeviceResponses(devices []entity.Device) []DeviceRs {
	logx.Debug("Invoke ToDeviceResponses")
	//goland:noinspection ALL
	deviceResponses := []DeviceRs{}
	for _, device := range devices {
		deviceResponses = append(deviceResponses, ToDeviceResponse(device))
	}
	return deviceResponses
}
