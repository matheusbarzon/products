package products

import (
	"encoding/json"
	"fmt"
	"net/http"
	entity "product/src/internal/category/entity"
	modelProduct "product/src/internal/category/models"

	"github.com/gorilla/mux"
)

func postProducts(response http.ResponseWriter, request *http.Request) {

	decoder := json.NewDecoder(request.Body)
	var product entity.Product
	err := decoder.Decode(&product)
	if err != nil {
		fmt.Fprintf(response, err.Error())
		return
	}

	modelProduct.Insert(&product)

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

	var structReturn entity.Product = modelProduct.SelectById(reqId)

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

	var structReturn = modelProduct.Select()

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
