package middleware

import (
	"net/http"

	"github.com/pedrocmart/canvas/internal/core"
)

// Logger print requests in console
func Logger(c *core.Config, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := core.GetCtxStringVal(r.Context(), core.ContextKeyRequestID)

		l := c.Log.Log.With().
			Str("url", r.URL.String()).
			Str("method", r.Method).
			Str("reqID", reqID).
			Logger()
		l.Info().Msg("request")
		h.ServeHTTP(w, r)
	})
}
