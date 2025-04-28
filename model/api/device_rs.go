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

import (
	"fmt"
	"github.com/pakaiwa/api/model/entity"
)

type DeviceRs struct {
	Id                 string `json:"id"`
	Status             string `json:"status"`
	PhoneNumber        string `json:"phone_number,omitempty"`
	CreatedAt          string `json:"created_at"`
	ConnectedAt        string `json:"connected_at,omitempty"`
	DisconnectedAt     string `json:"disconnected_at,omitempty"`
	DisconnectedReason string `json:"disconnected_reason,omitempty"`
}

func ToDeviceResponse(device entity.Device) DeviceRs {
	fmt.Println("Invoke ToDeviceResponse")
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
	fmt.Println("Invoke ToDeviceResponses")
	var deviceResponses []DeviceRs
	for _, device := range devices {
		deviceResponses = append(deviceResponses, ToDeviceResponse(device))
	}
	return deviceResponses
}
