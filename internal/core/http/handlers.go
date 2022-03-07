package http

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/pedrocmart/canvas/internal/core"
)

func NewHandler(c *core.Config, s core.Service) *Handler {
	return &Handler{
		config:  c,
		service: s,
	}
}

type Handler struct {
	config  *core.Config
	service core.Service
}

func (h *Handler) writeResponse(w http.ResponseWriter, r *http.Request, resp interface{}) {
	reqID := core.GetCtxStringVal(r.Context(), core.ContextKeyRequestID)
	start := core.GetCtxTimeVal(r.Context(), core.ContextKeyRequestTime)
	end := time.Now()

	l := h.config.Log.Log.With().
		Str("reqID", reqID).
		Str("duration", end.Sub(start).String()).
		Interface("response", resp).
		Logger()
	l.Info().Msg("response")

	bytes, err := json.Marshal(resp)
	if err != nil {
		h.writeError(w, r, ErrorWrapper{
			InnerError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			},
		})

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(bytes)
	if err != nil {
		log.Println(err)
	}
}

type ErrorWrapper struct {
	Error InnerError `json:"error" xml:"error"`
}

// InnerError this encodes the fields of the object
type InnerError struct {
	Code    int    `json:"code,omitempty" xml:"code,omitempty"`
	Message string `json:"message" xml:"message"`
	Status  int    `json:"status" xml:"status"`
}

func (h *Handler) writeError(w http.ResponseWriter, r *http.Request, er ErrorWrapper) {
	reqID := core.GetCtxStringVal(r.Context(), core.ContextKeyRequestID)
	start := core.GetCtxTimeVal(r.Context(), core.ContextKeyRequestTime)
	end := time.Now()

	l := h.config.Log.Log.With().
		Str("requestID", reqID).
		Str("duration", end.Sub(start).String()).
		Interface("response", er).
		Logger()
	l.Error().Msg("response")

	bytes, err := json.Marshal(er)
	if err != nil {
		log.Println("marshal", err)

		return
	}

	// write header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(er.Error.Status)

	// write return body
	_, err = w.Write(bytes)
	if err != nil {
		log.Println(err)
	}
}

func getParameter(values url.Values, key string) string {
	if len(values[key]) == 0 {
		return ""
	}

	return values[key][0]
}
