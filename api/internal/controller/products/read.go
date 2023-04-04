package products

import (
	"context"

	"github.com/Quoc-Bao-213/getgo-project/api/internal/repository/dbmodel"
)

// Create creates new product
func (i impl) GetAllProducts(ctx context.Context) (dbmodel.ProductSlice, error) {
	return i.repo.Inventory().GetAllProducts(ctx)
}
