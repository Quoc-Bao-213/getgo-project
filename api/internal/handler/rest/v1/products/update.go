package products

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Quoc-Bao-213/getgo-project/api/internal/controller/products"
	"github.com/Quoc-Bao-213/getgo-project/api/internal/httpserver"
	"github.com/go-chi/chi/v5"
)

// Update product
func (h Handler) Update() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()
		input := products.UpdateInput{}
		pid := chi.URLParam(r, "productID")
		ID, _ := strconv.ParseInt(pid, 0, 0)

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			return err
		}

		resp, err := h.productCtrl.Update(ctx, input, ID)
		if err != nil {
			return err
		}

		httpserver.RespondJSON(w, resp)

		return nil
	})
}
