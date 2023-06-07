package product

import (
	"errors"
	"testing"
	"time"

	product_domain "github.com/archel/product_store/pkg/product/domain"
	mocks "github.com/archel/product_store/test/mocks"
)

func fixedUuidGenerator() (string, error) {
	return "111-111-222", nil
}

func errorUuidGenerator() (string, error) {
	return "", errors.New("cannot generate id")
}

func TestCreatesAProduct(t *testing.T) {
	p := product_domain.NewProduct("111-111-111", 10.10, time.Now(), "Carrot Cake")
	pr := mocks.NewMockProductRepository(t)
	pr.EXPECT().Save(p).Return(nil)
	cp := NewCreateProductWithoutGeneratingId(pr)
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

func TestCannotStoreAProduct(t *testing.T) {
	p := product_domain.NewProduct("111-111-111", 10.10, time.Now(), "Carrot Cake")
	pr := mocks.NewMockProductRepository(t)
	pr.EXPECT().Save(p).Return(errors.New("test"))
	cp := NewCreateProductWithoutGeneratingId(pr)
	_, err := cp.Create(p)

	if err == nil {
		t.Errorf("Expected error when storing a product but none returned")
		return
	}
}

func TestCreatesAProductGeneratingTheId(t *testing.T) {
	now := time.Now()
	p := product_domain.NewProduct("", 10.10, now, "Carrot Cake")
	expected := product_domain.NewProduct("111-111-222", 10.10, now, "Carrot Cake")
	pr := mocks.NewMockProductRepository(t)
	pr.EXPECT().Save(expected).Return(nil)
	cp := NewCreateProductGeneratingId(pr, fixedUuidGenerator)
	got, err := cp.Create(p)

	if err != nil {
		t.Errorf("Got an error while creating a product %v", err)
		return
	}
	if got != expected {
		t.Errorf("Expected product to be %v but got %v", expected, got)
		return
	}

	pr.AssertCalled(t, "Save", expected)
}

func TestCannotGenerateProductId(t *testing.T) {
	p := product_domain.NewProduct("", 10.10, time.Now(), "Carrot Cake")
	pr := mocks.NewMockProductRepository(t)
	cp := NewCreateProductGeneratingId(pr, errorUuidGenerator)
	_, err := cp.Create(p)

	if err == nil {
		t.Errorf("Expected error when generating the id of the product but none returned")
		return
	}
}

func TestCannotStoreAProductWhenGeneratingTheId(t *testing.T) {
	now := time.Now()
	p := product_domain.NewProduct("", 10.10, now, "Carrot Cake")
	pWithId := product_domain.NewProduct("111-111-222", 10.10, now, "Carrot Cake")
	pr := mocks.NewMockProductRepository(t)
	pr.EXPECT().Save(pWithId).Return(errors.New("test"))
	cp := NewCreateProductGeneratingId(pr, fixedUuidGenerator)
	_, err := cp.Create(p)

	if err == nil {
		t.Errorf("Expected error when storing a product but none returned")
		return
	}
}
