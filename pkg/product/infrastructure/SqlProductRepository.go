package product

import (
	"fmt"

	product "github.com/archel/product_store/pkg/product/domain"
	"gorm.io/gorm"
)

type SqlProductRepository struct {
	conn *gorm.DB
}

func NewSqlProductRepository(conn *gorm.DB) SqlProductRepository {
	return SqlProductRepository{conn}
}

func (pr *SqlProductRepository) FindById(id string) (product.Product, error) {
	var p product.Product
	result := pr.conn.First(&p, "id = ?", id)
	if result.Error != nil {
		return p, fmt.Errorf("product not found by id %s", id)
	}
	return p, nil
}

func (pr *SqlProductRepository) Save(p product.Product) error {
	result := pr.conn.Save(&p)
	if result.Error != nil {
		return fmt.Errorf("cannot save product with id %s", p.Id)
	}
	return nil
}
