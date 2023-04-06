package products

import (
	"context"

	"github.com/Quoc-Bao-213/getgo-project/api/internal/model"
)

// UpdateInput represents for input to create Product
type UpdateInput struct {
	Name        string
	Description string
	Price       int64
}

// Update product
func (i impl) Update(ctx context.Context, input UpdateInput, pid int64) (model.Product, error) {
	return i.repo.Inventory().UpdateProduct(ctx, model.Product{
		Price:       input.Price,
		Description: input.Description,
		Name:        input.Name,
		Status:      model.ProductStatusActive,
	}, pid)
}
