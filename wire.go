package main

import (
	"github.com/nilerajput91/gingormframework/product"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func initProductAPI(db *gorm.DB) product.ProductAPI {
	wire.Build(product.ProvideProductRepostiory, product.ProvideProductService, product.ProvideProductAPI)

	return product.ProductAPI{}
}
