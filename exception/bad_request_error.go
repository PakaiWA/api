// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Thu 01/05/25 19.15
// @project api exception
//

package exception

type BadRequestError struct {
	Error string
}

func NewBadRequestError(error string) BadRequestError {
	return BadRequestError{Error: error}
}
