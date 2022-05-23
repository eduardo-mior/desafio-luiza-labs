package auth_controller

import (
	"desafio_luizalabs/domain/mocks"
	"desafio_luizalabs/viewmodels"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

func TestGerarTokenJWTSucesso(t *testing.T) {

	serviceAuth := new(mocks.IServiceAuth)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)

	serviceAuth.On("GenerateTokenJWT").Return("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UifQ.b-bQWDLukCsP-kz255shgiDNoPjycxSizc1cIljoEic", nil)
	GerarTokenJWT(serviceAuth)(c)

	responseBody := viewmodels.Token{}
	err := json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, nil, err)

	assert.Equal(t, 200, response.Code)
	assert.NotEmpty(t, responseBody.Token)

}

func TestGerarTokenJWTErro(t *testing.T) {

	serviceAuth := new(mocks.IServiceAuth)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)

	serviceAuth.On("GenerateTokenJWT").Return("", jwt.ErrInvalidKeyType)
	GerarTokenJWT(serviceAuth)(c)

	responseBody := viewmodels.Error{}
	err := json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, nil, err)

	assert.Equal(t, 500, response.Code)
	assert.Equal(t, viewmodels.Error{Error: responseBody.Error}, responseBody)

}
