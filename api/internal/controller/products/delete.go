package products

import (
	"context"
)

// Delete product
func (i impl) Delete(ctx context.Context, pid int64) (bool, error) {
	return i.repo.Inventory().DeleteProduct(ctx, pid)
}
