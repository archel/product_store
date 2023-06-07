package product

import (
	"errors"

	product "github.com/archel/product_store/pkg/product/domain"
)

type CreateProduct interface {
	Create(p product.Product) (product.Product, error)
}

type CreateProductWithoutGeneratingId struct {
	r product.ProductRepository
}

func (cp CreateProductWithoutGeneratingId) Create(p product.Product) (product.Product, error) {
	if err := cp.r.Save(p); err != nil {
		return product.Product{}, errors.New("cannot create a product")
	}

	return p, nil
}

type CreateProductGeneratingId struct {
	g  product.ProductIdGenerator
	cp CreateProduct
}

func (cp CreateProductGeneratingId) Create(p product.Product) (product.Product, error) {
	err := p.GenerateId(cp.g)
	if err != nil {
		return product.Product{}, err
	}

	if p, err := cp.cp.Create(p); err != nil {
		return product.Product{}, errors.New("cannot create a product")
	} else {
		return p, nil
	}

}

func NewCreateProductWithoutGeneratingId(r product.ProductRepository) CreateProduct {
	return CreateProductWithoutGeneratingId{
		r: r,
	}
}

func NewCreateProductGeneratingId(r product.ProductRepository, g product.ProductIdGenerator) CreateProduct {
	return CreateProductGeneratingId{
		cp: NewCreateProductWithoutGeneratingId(r),
		g:  g,
	}
}
