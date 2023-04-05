package inventory

import (
	"context"
	"database/sql"

	"github.com/Quoc-Bao-213/getgo-project/api/internal/model"
	"github.com/Quoc-Bao-213/getgo-project/api/internal/repository/dbmodel"
)

type Repository interface {
	// specification
	CreateProduct(context.Context, model.Product) (model.Product, error)
	GetAllProducts(context.Context) (dbmodel.ProductSlice, error)
	GetProductDetails(context.Context, int64) (model.Product, error)
	DeleteProduct(context.Context, int64) (bool, error)
	UpdateProduct(context.Context, model.Product, int64) (model.Product, error)
}

func New(db *sql.DB) Repository {
	return impl{
		db: db,
	}
}

type impl struct {
	db *sql.DB
}
