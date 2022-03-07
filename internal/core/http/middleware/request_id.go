package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/pedrocmart/canvas/internal/core"
)

// ReqID add requestID to request
func ReqID(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, core.ContextKeyRequestID, uuid.New().String())
		ctx = context.WithValue(ctx, core.ContextKeyRequestTime, time.Now())
		r = r.WithContext(ctx)

		h.ServeHTTP(w, r)
	})
}
