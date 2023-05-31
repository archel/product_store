package product

import "time"

type Product struct {
	Id          string
	Price       float64
	Creation    time.Time
	Description string
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
