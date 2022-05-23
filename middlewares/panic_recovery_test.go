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

func TestPanicRecovery(t *testing.T) {

	router := gin.Default()
	router.Use(PanicRecovery())
	router.GET("/panic-recovery", func(c *gin.Context) {
		panic("teste")
	})

	server := httptest.NewServer(router)
	defer server.Close()

	request, err := http.NewRequest("GET", "http://localhost:8080/panic-recovery", nil)
	assert.Equal(t, nil, err)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	responseBody := viewmodels.Error{}
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, nil, err)

	assert.Equal(t, 500, response.Code)
	assert.Equal(t, viewmodels.Error{Error: Erro500Interno}, responseBody)

}
