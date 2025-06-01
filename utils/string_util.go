// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Mon 28/04/25 15.46
// @project api https://github.com/PakaiWA/api/tree/main/utils
//

package utils

import (
	"database/sql"

	"github.com/google/uuid"
)

func SafeString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

func GenerateUUID() string {
	return "pwa-" + uuid.New().String()
}
