package middlewares

import (
	"desafio_luizalabs/viewmodels"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMiddlewareNotFound(t *testing.T) {

	router := gin.Default()
	router.NoRoute(NotFound())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})

	server := httptest.NewServer(router)
	defer server.Close()

	request, err := http.NewRequest("GET", "http://localhost:8080/end-point-inexistente ", nil)
	assert.Equal(t, nil, err)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	responseBody := viewmodels.Error{}
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, nil, err)

	assert.Equal(t, 404, response.Code)
	assert.Equal(t, viewmodels.Error{Error: Erro404RecursoNaoEncontrado}, responseBody)

}
