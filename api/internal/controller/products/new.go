package products

import (
	"context"

	"github.com/Quoc-Bao-213/getgo-project/api/internal/model"
	"github.com/Quoc-Bao-213/getgo-project/api/internal/repository"
	"github.com/Quoc-Bao-213/getgo-project/api/internal/repository/dbmodel"
)

// Controller represents the specification of this pkg
type Controller interface {
	Create(context.Context, CreateInput) (model.Product, error)
	GetAllProducts(context.Context) (dbmodel.ProductSlice, error)
	GetProductDetails(context.Context, int64) (model.Product, error)
	Delete(context.Context, int64) (bool, error)
}

// New returns an implementation instance which satisfying Controller
func New(repo repository.Registry) Controller {
	return impl{
		repo: repo,
	}
}

type impl struct {
	repo repository.Registry
}
