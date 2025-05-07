// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sun 27/04/25 22.01
// @project api https://github.com/PakaiWA/api/tree/main/helper
//

package helper

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func DBTransaction(ctx context.Context, DB *pgxpool.Pool) (pgx.Tx, *pgxpool.Conn, error) {
	fmt.Println("Invoke DBTransaction")
	conn, err := DB.Acquire(ctx)
	if err != nil {
		return nil, nil, err
	}

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		conn.Release()
		return nil, nil, err
	}

	return tx, conn, nil
}

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
