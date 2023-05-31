package product

import (
	"errors"

	product "github.com/archel/product_store/pkg/product/domain"
)

type CreateProduct struct {
	r product.ProductRepository
}

func NewCreateProduct(r product.ProductRepository) CreateProduct {
	return CreateProduct{
		r: r,
	}
}

func (cp CreateProduct) Create(p product.Product) (product.Product, error) {
	if err := cp.r.Save(p); err != nil {
		return product.Product{}, errors.New("cannot create a product")
	}

	return p, nil
}
