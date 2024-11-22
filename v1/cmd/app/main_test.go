package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

// Test the HomePage response
func TestHomePageHandler(t *testing.T) {

	mockResponse := `{"message":"This is the homepage"}`

	r := SetUpRouter()
	r.GET("/", PingHandler)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
