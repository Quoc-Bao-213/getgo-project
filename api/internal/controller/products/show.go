package products

import (
	"context"

	"github.com/Quoc-Bao-213/getgo-project/api/internal/model"
)

// Create creates new product
func (i impl) GetProductDetails(ctx context.Context, pid int64) (model.Product, error) {
	return i.repo.Inventory().GetProductDetails(ctx, pid)
}
