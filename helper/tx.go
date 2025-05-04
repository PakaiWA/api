// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sun 27/04/25 22.01
// @project api helper
//

package helper

import (
	"context"
	"github.com/jackc/pgx/v5"
)

func CommitOrRollback(ctx context.Context, tx pgx.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback(ctx)
		PanicIfError(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit(ctx)
		PanicIfError(errorCommit)
	}
}
