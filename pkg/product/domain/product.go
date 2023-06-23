package product

import (
	"errors"
	"time"
)

type Product struct {
	Id          string    `json:"id,omitempty" uri:"id" gorm:"primaryKey,column=id"`
	Price       float64   `json:"price" gorm:"column=price"`
	Creation    time.Time `json:"creation" gorm:"column=creation"`
	Description string    `json:"description" gorm:"column=description"`
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
	FindById(id string) (Product, error)
	Save(p Product) error
}

type ProductIdGenerator func() (string, error)
