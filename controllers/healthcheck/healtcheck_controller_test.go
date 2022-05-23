package healthcheck_controller

import (
	"desafio_luizalabs/viewmodels"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	HealthCheck()(c)

	responseBody := viewmodels.HealthCheck{}
	err := json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, nil, err)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, viewmodels.HealthCheck{Status: AvailableStatus}, responseBody)

}
