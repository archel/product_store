package product

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	product "github.com/archel/product_store/pkg/product/domain"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/product", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func TestCreatesAProduct(t *testing.T) {
	t.Skip("disabled")
	router := setupRouter()
	now := time.Now()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/product", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status to be %d but got %d", http.StatusCreated, w.Code)
	}

	body := w.Body.String()
	p := product.Product{}
	err := json.Unmarshal([]byte(body), &p)
	ep := product.NewProduct("123123", 10.10, now, "Carrot Cake")
	if err != nil || p == ep {
		t.Errorf("Expected body to be %v but got %v", ep, p)
	}
}
