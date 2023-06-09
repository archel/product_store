package product

import (
	"net/http"

	product_application "github.com/archel/product_store/pkg/product/application"
	product_domain "github.com/archel/product_store/pkg/product/domain"
	product_infra "github.com/archel/product_store/pkg/product/infrastructure"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductController struct {
	cp                 product_application.CreateProduct
	cpWithIdGeneration product_application.CreateProduct
}

func NewProductController(cp product_application.CreateProduct, cpWithIdGeneration product_application.CreateProduct) ProductController {
	return ProductController{cp: cp, cpWithIdGeneration: cpWithIdGeneration}
}

func (c ProductController) CreateProductHandler(ctx *gin.Context) {
	var p product_domain.Product
	ctx.BindJSON(p)
	id := ctx.Param("id")
	var generator = func() (string, error) { return id, nil }
	p.GenerateId(generator)

	p, err := c.cp.Create(p)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, p)
}

func (c ProductController) CreateProductGeneratingIdHandler(ctx *gin.Context) {
	var p product_domain.Product
	ctx.BindJSON(p)

	p, err := c.cpWithIdGeneration.Create(p)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, p)
}

func SetupProductRoutes(r *gin.Engine, conn *gorm.DB) {
	repo := product_infra.NewSqlProductRepository(conn)
	cp := product_application.NewCreateProductWithoutGeneratingId(&repo)
	cpWithId := product_application.NewCreateProductGeneratingId(&repo, product_infra.UuidProductGenerator)
	c := NewProductController(cp, cpWithId)
	r.POST("/product/:id", c.CreateProductHandler)
	r.POST("/product/", c.CreateProductGeneratingIdHandler)
}
