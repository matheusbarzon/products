package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	port := "4000"

	message := fmt.Sprintf("API is running on port %s!", port)

	http.HandleFunc(
		"/",
		func(res http.ResponseWriter, req *http.Request) {

			responseMap := map[string]string{"itsRunning": message}

			json, err := json.Marshal(responseMap)

			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}

			res.WriteHeader(200)
			res.Header().Set("Content-Type", "application/json")
			res.Write(json)
		},
	)

	log.Println(message)
	http.ListenAndServe(":"+port, nil)
}
