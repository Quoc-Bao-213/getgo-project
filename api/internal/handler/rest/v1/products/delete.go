package products

import (
	"net/http"
	"strconv"

	"github.com/Quoc-Bao-213/getgo-project/api/internal/httpserver"
	"github.com/go-chi/chi/v5"
)

func (h Handler) Delete() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()
		pid := chi.URLParam(r, "productID")

		ID, err := strconv.ParseInt(pid, 0, 0)

		if err != nil {
			return &httpserver.Error{
				Status: http.StatusBadRequest,
				Code: "bad_request",
				Desc: "Invalid product ID",
			}
		}

		resp, err := h.productCtrl.Delete(ctx, ID)

		if err != nil {
			return err
		}

		var obj = map[string]interface{}{}

		if resp {
			obj["status"] = "Delete product successfully"
		}

		httpserver.RespondJSON(w, obj)

		return nil
	})
}


