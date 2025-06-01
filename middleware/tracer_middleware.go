// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Sun 01/06/25 16.38
// @project api middleware
// https://github.com/PakaiWA/api/tree/main/middleware
//

package middleware

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"github.com/pakaiwa/api/utils"
	"net/http"
)

func TraceIdMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		traceID := r.Header.Get("ax-request-id")
		if traceID == "" {
			traceID = utils.GenerateUUID()
		}
		ctx := context.WithValue(r.Context(), "trace_id", traceID)
		next(w, r.WithContext(ctx), ps)
	}
}
