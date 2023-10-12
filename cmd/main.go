package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	products "product/pkg/handlers"

	"github.com/gorilla/mux"
)

func main() {
	var port string = getenv("PORT", "4000")
	message := fmt.Sprintf("API is running on port %s!", port)

	handlers := mux.NewRouter()

	products.HandleRequests(handlers)

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
