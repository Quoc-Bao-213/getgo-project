package products

import (
	"net/http"

	"github.com/Quoc-Bao-213/getgo-project/api/internal/httpserver"
)

// Show all products
func (h Handler) GetAllProducts() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		resp, err := h.productCtrl.GetAllProducts(ctx)
		if err != nil {
			return err
		}

		httpserver.RespondJSON(w, resp)

		return nil
	})
}
