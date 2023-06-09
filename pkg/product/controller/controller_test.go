package product

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	product_domain "github.com/archel/product_store/pkg/product/domain"
	"github.com/archel/product_store/test/containers"
	"github.com/gin-gonic/gin"
)

func setupRouter(c ProductController) *gin.Engine {
	r := gin.Default()
	r.POST("/product", c.CreateProductHandler)
	return r
}

func setupController(t *testing.T) ProductController {
	// cp := mocks.NewMockCreateProduct(t)
	// cpWithIdGeneration := mocks.NewMockCreateProduct(t)
	// return NewProductController(cp, cpWithIdGeneration)
	return ProductController{}
}

func TestCreatesAProduct(t *testing.T) {
	container, _ := containers.NewPostgresContainer()
	defer container.Terminate()

	c := setupController(t)
	router := setupRouter(c)
	now := time.Now()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/product", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status to be %d but got %d", http.StatusCreated, w.Code)
	}

	body := w.Body.String()
	var p product_domain.Product
	json.Unmarshal([]byte(body), &p)
	ep := product_domain.NewProduct("123123", 10.10, now, "Carrot Cake")
	if p == ep {
		t.Errorf("Expected body to be %v but got %v", ep, p)
	}
}
