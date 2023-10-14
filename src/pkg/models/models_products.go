package models_products

import (
	"log"
	"product/src/pkg/database"
	structs_products "product/src/pkg/structs"
	"strconv"
)

func Insert(product *structs_products.Product) {

	stament, err := database.Connection().Prepare(
		`insert into products
			(nome, gtin, inclusao)
		values
			(?, ?, ?)`)

	if err != nil {
		log.Printf(err.Error())
		return
	}

	defer stament.Close()

	result, err := stament.Execute(product.Nome, product.Gtin, product.Inclusao)

	if err != nil {
		log.Printf(err.Error())
		return
	}

	product.Id = int64(result.InsertId)

	defer result.Close()
}

func SelectById(reqId string) structs_products.Product {

	stament, err := database.Connection().Prepare(
		`select id, nome, gtin, inclusao
		   from products
		  where id = ?`)

	if err != nil {
		log.Printf(err.Error())
		// return
	}

	defer stament.Close()

	result, err := stament.Execute(reqId)

	if err != nil {
		log.Printf(err.Error())
		// return
	}

	var structReturn structs_products.Product

	id, _ := strconv.ParseInt(reqId, 10, 64)
	nome, _ := result.GetStringByName(0, "nome")
	gtin, _ := result.GetStringByName(0, "gtin")
	inclusao, _ := result.GetStringByName(0, "inclusao")

	structReturn = structs_products.Product{
		Id:       id,
		Nome:     nome,
		Gtin:     gtin,
		Inclusao: inclusao,
	}

	defer result.Close()

	return structReturn
}

func Select() structs_products.Products {

	result, err := database.Connection().Execute(
		`select id, nome, gtin, inclusao
		   from products`)

	if err != nil {
		log.Printf(err.Error())
		// return
	}

	var structReturn structs_products.Products

	for i := range result.Values {

		id, _ := result.GetIntByName(i, "id")
		nome, _ := result.GetStringByName(i, "nome")
		gtin, _ := result.GetStringByName(i, "gtin")
		inclusao, _ := result.GetStringByName(i, "inclusao")

		structReturn = append(structReturn, structs_products.Product{
			Id:       id,
			Nome:     nome,
			Gtin:     gtin,
			Inclusao: inclusao,
		})
	}

	defer result.Close()

	return structReturn
}
