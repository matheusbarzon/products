package products

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func PostProducts(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Produto criado com sucesso!")
}

func GetByIdProducts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := vars["id"]
	if err {
		fmt.Println("id is missing in parameters")
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Listagem do produto com `id` %s!", id)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Listagem de produtos!")
}
