// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sun 27/04/25 17.11
// @project api https://github.com/PakaiWA/api/tree/main/helper

package helper

import "fmt"

func PanicIfError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
