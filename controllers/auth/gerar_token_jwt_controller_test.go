package auth_controller

import (
	"desafio_luizalabs/services"
	"desafio_luizalabs/viewmodels"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGerarTokenJWT(t *testing.T) {

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	GerarTokenJWT(&services.ServiceAuth{})(c)

	responseBody := viewmodels.Token{}
	err := json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, nil, err)

	assert.Equal(t, 200, response.Code)
	assert.NotEmpty(t, responseBody.Token)

}
