package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	products "product/pkg/handlers"

	"github.com/gorilla/mux"
)

func statusApi(w http.ResponseWriter, r *http.Request) {
	var port string = getenv("PORT", "4000")
	message := fmt.Sprintf("API is running on port %s!", port)

	responseMap := map[string]string{"itsRunning": message}

	json, err := json.Marshal(responseMap)

	if err != nil {
		fmt.Fprintf(w, "Error encoding JSON: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", json)
}

func handleRequests(handlers *mux.Router) {

	handlers.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {

		methods := map[string]func(w http.ResponseWriter, r *http.Request){}

		methods[http.MethodGet] = products.GetProducts
		methods[http.MethodPost] = products.PostProducts

		methods[r.Method](w, r)

	})

	handlers.HandleFunc("/products/{id}", products.GetByIdProducts)

	http.HandleFunc("/", statusApi)
}

func main() {
	var port string = getenv("PORT", "4000")
	message := fmt.Sprintf("API is running on port %s!", port)

	handlers := mux.NewRouter()

	handleRequests(handlers)

	log.Println(message)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), handlers))
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
