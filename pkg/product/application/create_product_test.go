package product

import (
	"testing"
	"time"

	product_domain "github.com/archel/product_store/pkg/product/domain"
	mocks "github.com/archel/product_store/test/mocks"
)

func TestCreatesAProduct(t *testing.T) {
	p := product_domain.NewProduct("111-111-111", 10.10, time.Now(), "Carrot Cake")
	pr := mocks.NewMockProductRepository(t)
	pr.EXPECT().Save(p).Return(nil)
	cp := NewCreateProduct(pr)
	got, err := cp.Create(p)

	if err != nil {
		t.Errorf("Got an error while creating a product %v", err)
		return
	}
	if got != p {
		t.Errorf("Expected product to be %v but got %v", p, got)
		return
	}

	pr.AssertCalled(t, "Save", p)
}
