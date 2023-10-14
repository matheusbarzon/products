package handlers_products

import (
	"encoding/json"
	"fmt"
	"net/http"
	models_products "product/src/pkg/models"
	structs_products "product/src/pkg/structs"

	"github.com/gorilla/mux"
)

func postProducts(response http.ResponseWriter, request *http.Request) {

	decoder := json.NewDecoder(request.Body)
	var product structs_products.Product
	err := decoder.Decode(&product)
	if err != nil {
		fmt.Fprintf(response, err.Error())
		return
	}

	models_products.Insert(&product)

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

	var structReturn structs_products.Product = models_products.SelectById(reqId)

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

	var structReturn = models_products.Select()

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
