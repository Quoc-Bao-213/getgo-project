package products

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Quoc-Bao-213/getgo-project/api/internal/controller/products"
	"github.com/Quoc-Bao-213/getgo-project/api/internal/httpserver"
)

// Create creates new product
func (h Handler) Create() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		input := products.CreateInput{}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			return err
		}

		if err := validate(input); err != nil {
			return err
		}

		resp, err := h.productCtrl.Create(ctx, input)
		if err != nil {
			return err
		}

		httpserver.RespondJSON(w, resp)

		return nil
	})
}

func validate(input products.CreateInput) error {
	if strings.TrimSpace(input.Name) == "" {
		return errInvalidName
	}

	if strings.TrimSpace(input.Description) == "" {
		return errInvalidDescription
	}

	if input.Price <= 0 {
		return errInvalidPrice
	}

	return nil
}

// func validateAll(input products.CreateInput) []error {
// 	var errs []error
// 	if strings.TrimSpace(input.Name) == "" {
// 		errs = append(errs, errInvalidName)
// 	}

// 	if strings.TrimSpace(input.Description) == "" {
// 		errs = append(errs, errInvalidDescription)
// 	}

// 	if input.Price <= 0 {
// 		errs = append(errs, errInvalidPrice)
// 	}

// 	if len(errs) > 0 {
// 		return errs
// 	}

// 	return nil
// }
