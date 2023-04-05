package products

import (
	"context"
)

func (i impl) Delete(ctx context.Context, productID int64) (bool, error) {
	return i.repo.Inventory().DeleteProduct(ctx, productID)
}
