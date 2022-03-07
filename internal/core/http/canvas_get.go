package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pedrocmart/canvas/internal/core"
	"github.com/pedrocmart/canvas/internal/core/errors"
)

func (h *Handler) CanvasGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	params := r.URL.Query()

	resp, err := h.service.CanvasGet(r.Context(),
		&core.CanvasGetRequest{
			ID: getParameter(params, "id"),
		},
	)
	if err != nil {
		e, ok := err.(errors.Error)
		if ok {
			h.writeError(w, r, ErrorWrapper{
				InnerError{
					Code:    e.Code,
					Message: e.Error(),
					Status:  e.Status,
				},
			})

			return
		}

		h.writeError(w, r, ErrorWrapper{
			InnerError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			},
		})

		return
	}

	h.writeResponse(w, r, resp)
}
