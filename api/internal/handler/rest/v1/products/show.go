package products

import (
	"net/http"
	"strconv"

	"github.com/Quoc-Bao-213/getgo-project/api/internal/httpserver"
	"github.com/go-chi/chi/v5"
)

func (h Handler) GetProductDetails() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()
		productID := chi.URLParam(r, "productID")

		ID, err := strconv.ParseInt(productID, 0, 0)

		if err != nil {
			return &httpserver.Error{
				Status: http.StatusBadRequest,
				Code: "bad_request",
				Desc: "Invalid product ID",
			}
		}

		resp, err := h.productCtrl.GetProductDetails(ctx, ID)
		if err != nil {
			return err
		}

		httpserver.RespondJSON(w, resp)

		return nil
	})
}
