package inventory

import (
	"context"

	"github.com/Quoc-Bao-213/getgo-project/api/internal/model"
	"github.com/Quoc-Bao-213/getgo-project/api/internal/repository/dbmodel"
	"github.com/Quoc-Bao-213/getgo-project/api/internal/repository/generator"
	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (i impl) CreateProduct(ctx context.Context, m model.Product) (model.Product, error) {
	id, err := generator.ProductSNF.Generate()
	if err != nil {
		return model.Product{}, err
	}

	o := dbmodel.Product{
		ID:          id,
		ExternalID:  m.ExternalID,
		Name:        m.Name,
		Description: m.Description,
		Status:      m.Status.String(),
		Price:       m.Price,
	}

	if err := o.Insert(ctx, i.db, boil.Infer()); err != nil {
		return model.Product{}, pkgerrors.WithStack(err)
	}

	m.ID = id
	m.CreatedAt = o.CreatedAt
	m.UpdatedAt = o.UpdatedAt

	return m, nil
}

func (i impl) GetAllProducts(ctx context.Context) (dbmodel.ProductSlice, error) {
	o, err := dbmodel.Products().All(ctx, i.db)

	return o, err
}

func (i impl) GetProductDetails(ctx context.Context, productID int64) (model.Product, error) {
	product, err := dbmodel.Products(dbmodel.ProductWhere.ID.EQ(productID)).One(ctx, i.db)

	if err != nil {
		return model.Product{}, pkgerrors.WithStack(err)
	}

	return model.Product{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Status:      model.ProductStatus(product.Status),
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}, nil
}
