package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	products "product/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.POST("/pessoas", products.PostProducts)
	router.GET("/pessoas/:id", products.GetByIdProducts)
	router.GET("/pessoas", products.GetProducts)

	var port string = getenv("PORT", "4000")
	message := fmt.Sprintf("API is running on port %s!", port)

	router.Run(fmt.Sprintf("localhost:%s", port))
	log.Println(message)

	router.GET(
		"/",
		func(c *gin.Context) {
			responseMap := map[string]string{"itsRunning": message}

			json, err := json.Marshal(responseMap)

			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, err.Error())
				return
			}

			c.IndentedJSON(http.StatusOK, json)
		})
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
