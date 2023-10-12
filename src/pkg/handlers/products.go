package products

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"product/pkg/database"
	"strconv"

	"github.com/gorilla/mux"
)

type Product struct {
	Id       int64  `json:"id"`
	Nome     string `json:"nome"`
	Gtin     string `json:"gtin"`
	Inclusao string `json:"inclusao"`
}

type Products []Product

func postProducts(response http.ResponseWriter, request *http.Request) {

	decoder := json.NewDecoder(request.Body)
	var product Product
	err := decoder.Decode(&product)
	if err != nil {
		fmt.Fprintf(response, err.Error())
		return
	}

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

	println(result.InsertId)

	jsonReturn, _ := json.Marshal(product)

	response.WriteHeader(http.StatusCreated)
	fmt.Fprintf(response, string(jsonReturn))
}

func getByIdProducts(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	reqId, reqErr := vars["id"]

	if !reqErr {
		fmt.Println("id is missing in parameters")
		return
	}

	stament, err := database.Connection().Prepare(
		`select id, nome, gtin, inclusao
		   from products
		  where id = ?`)

	if err != nil {
		log.Printf(err.Error())
		return
	}

	defer stament.Close()

	result, err := stament.Execute(reqId)

	if err != nil {
		log.Printf(err.Error())
		return
	}

	var structReturn Product

	id, _ := strconv.ParseInt(reqId, 10, 64)
	nome, _ := result.GetStringByName(0, "nome")
	gtin, _ := result.GetStringByName(0, "gtin")
	inclusao, _ := result.GetStringByName(0, "inclusao")

	structReturn = Product{
		Id:       id,
		Nome:     nome,
		Gtin:     gtin,
		Inclusao: inclusao,
	}

	defer result.Close()

	jsonReturn, err := json.Marshal(structReturn)
	if err != nil {
		fmt.Println(err)
		return
	}

	response.WriteHeader(http.StatusOK)

	fmt.Fprintf(response, string(jsonReturn))
}

func getProducts(response http.ResponseWriter, _ *http.Request) {
	response.WriteHeader(http.StatusOK)

	database.Connection()

	result, err := database.Connection().Execute(
		`select id, nome, gtin, inclusao
		   from products`)

	if err != nil {
		log.Printf(err.Error())
		return
	}

	var structReturn Products

	for i := range result.Values {

		id, _ := result.GetIntByName(i, "id")
		nome, _ := result.GetStringByName(i, "nome")
		gtin, _ := result.GetStringByName(i, "gtin")
		inclusao, _ := result.GetStringByName(i, "inclusao")

		structReturn = append(structReturn, Product{
			Id:       id,
			Nome:     nome,
			Gtin:     gtin,
			Inclusao: inclusao,
		})
	}

	defer result.Close()

	jsonReturn, err := json.Marshal(structReturn)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(response, string(jsonReturn))
}

func HandleRequests(handlers *mux.Router) {

	handlers.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {

		methods := map[string]func(w http.ResponseWriter, r *http.Request){}

		methods[http.MethodGet] = getProducts
		methods[http.MethodPost] = postProducts

		methods[r.Method](w, r)

	})

	handlers.HandleFunc("/products/{id}", getByIdProducts)
}
