package products

import (
	"context"

	"github.com/Quoc-Bao-213/getgo-project/api/internal/repository/dbmodel"
)

// Show all products
func (i impl) GetAllProducts(ctx context.Context) (dbmodel.ProductSlice, error) {
	return i.repo.Inventory().GetAllProducts(ctx)
}
