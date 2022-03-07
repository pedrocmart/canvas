package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pedrocmart/canvas/internal/core"
	"github.com/pedrocmart/canvas/internal/core/errors"
)

func (h *Handler) CanvasCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.writeError(w, r, ErrorWrapper{
			InnerError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			},
		})

		return
	}

	var req core.CanvasCreateRequest

	err = json.Unmarshal(body, &req)
	if err != nil {
		h.writeError(w, r, ErrorWrapper{
			InnerError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			},
		})

		return
	}

	resp, err := h.service.CanvasCreate(r.Context(), &req)
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
