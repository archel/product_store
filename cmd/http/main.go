package main

import (
	"fmt"
	"os"

	product_controller "github.com/archel/product_store/pkg/product/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	r := gin.New()
	dsn := "host=localhost user=postgres password=s3cr3t dbname=postgres port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error stabilishing connection to the db %v", err)
		os.Exit(-1)
	}
	product_controller.SetupProductRoutes(r, db)
	r.Run(":8080")
}
