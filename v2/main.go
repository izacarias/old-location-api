package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.GET("/", getRoot)

	route.Run(":8080")
}

func getRoot(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}
