package product

import (
	"errors"
	"time"
)

type Product struct {
	Id          string
	Price       float64
	Creation    time.Time
	Description string
}

func (p *Product) GenerateId(generator ProductIdGenerator) error {
	if id, err := generator(); err != nil {
		return errors.New("cannot generate a valid id for this product")
	} else {
		p.Id = id
	}
	return nil
}

func NewProduct(
	id string,
	price float64,
	creation time.Time,
	description string,
) Product {
	return Product{
		Id:          id,
		Price:       price,
		Creation:    creation,
		Description: description,
	}
}

type ProductRepository interface {
	FindById() Product
	Save(p Product) error
}

type ProductIdGenerator func() (string, error)
