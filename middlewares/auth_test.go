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

func TestMiddlewareAuth(t *testing.T) {

	router := gin.Default()
	router.Use(AuthJWT())
	router.GET("/end-point-privado", func(c *gin.Context) {
		c.Status(200)
	})

	server := httptest.NewServer(router)
	defer server.Close()

	request, err := http.NewRequest("GET", "http://localhost:8080/end-point-privado", nil)
	assert.Equal(t, nil, err)

	// Erro de token não informado
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	responseBody := viewmodels.AuthError{}
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, nil, err)

	assert.Equal(t, 401, response.Code)
	assert.Equal(t, viewmodels.AuthError{Error: erroTokenNaoInformado}, responseBody)

	// Erro de token invalido (sem a palavra Bearer na frente)
	request.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjF9.kwiyVA1maOQtr-rrH5dIAem6SJgYckUtWZcyBDTvLsU")
	assert.Equal(t, nil, err)

	response2 := httptest.NewRecorder()
	router.ServeHTTP(response2, request)

	responseBody2 := viewmodels.AuthError{}
	err = json.NewDecoder(response2.Body).Decode(&responseBody2)
	assert.Equal(t, nil, err)

	assert.Equal(t, 401, response2.Code)
	assert.Equal(t, viewmodels.AuthError{Error: erroTokenInvalido}, responseBody2)

	// Erro de token invalido (secret usada para gerar o token inválida)
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjF9.kwiyVA1maOQtr-rrH5dIAem6SJgYckUtWZcyBDTvLsU")
	assert.Equal(t, nil, err)

	response3 := httptest.NewRecorder()
	router.ServeHTTP(response3, request)

	responseBody3 := viewmodels.AuthError{}
	err = json.NewDecoder(response3.Body).Decode(&responseBody3)
	assert.Equal(t, nil, err)

	assert.Equal(t, 401, response3.Code)
	assert.Equal(t, viewmodels.AuthError{Error: erroTokenInvalido}, responseBody3)

	// Token valido
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjF9.g8-pxhGP26OP5eZxjg2Df3RpXwQvPeHUFabYya28pBk")
	assert.Equal(t, nil, err)

	response4 := httptest.NewRecorder()
	router.ServeHTTP(response4, request)

	assert.Equal(t, 200, response4.Code)

}
