package models_products

import (
	"product/src/pkg/database"
	structs_products "product/src/pkg/structs"
)

func Insert(product *structs_products.Product) {

	db := database.Connection()

	db.Save(&product)
}

func SelectById(reqId string) structs_products.Product {

	db := database.Connection()

	var structReturn structs_products.Product

	db.First(&structReturn, reqId)

	return structReturn
}

func Select() structs_products.Products {

	db := database.Connection()

	var structReturn structs_products.Products

	db.Find(&structReturn)

	return structReturn
}
