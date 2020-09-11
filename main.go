package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/nilerajput91/gingormframework/product"
)

func initDB() *gorm.DB {
	dataSourceName := "root:root@tcp(localhost:3306)/nilesraj?parseTime=True"
	db, err := gorm.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&product.Product{})

	return db

}

func main() {
	db := initDB()
	defer db.Close()

	productAPI := initProductAPI(db)

	r := gin.Default()

	r.GET("/products", productAPI.FindAll)
	r.GET("/products/:id", productAPI.FindByID)
	r.POST("/products", productAPI.Create)
	r.PUT("/products/:id", productAPI.Update)
	r.DELETE("/products/:id", productAPI.Delete)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
