package products

import (
	entity "product/src/internal/category/entity"
	"product/src/pkg/database"
)

func Insert(product *entity.Product) {

	db := database.Connection()

	db.Save(&product)
}

func SelectById(reqId string) entity.Product {

	db := database.Connection()

	var structReturn entity.Product

	db.First(&structReturn, reqId)

	return structReturn
}

func Select() entity.Products {

	db := database.Connection()

	var structReturn entity.Products

	db.Find(&structReturn)

	return structReturn
}
