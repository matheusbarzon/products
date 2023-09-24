package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusCreated, "")
}

func GetByIdProducts(c *gin.Context) {
	id := c.Param("id")

	c.IndentedJSON(http.StatusOK, id)
}

func GetProducts(c *gin.Context) {
	qryParam := c.Query("t")

	c.IndentedJSON(http.StatusOK, qryParam)
}
