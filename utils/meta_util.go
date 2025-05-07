// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Mon 28/04/25 10.55
// @project api https://github.com/PakaiWA/api/tree/main/utils
//

package utils

import (
	"fmt"
	"net/http"
)

func GetMetaLocation(r *http.Request) string {
	return fmt.Sprintf("%s%s/", GetHost(r), r.RequestURI)
}

func GetHost(r *http.Request) string {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	return fmt.Sprintf("%s://%s", scheme, r.Host)

}
