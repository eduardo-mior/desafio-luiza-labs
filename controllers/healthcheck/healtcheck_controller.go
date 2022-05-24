package healthcheck_controller

import (
	"desafio_luizalabs/viewmodels"

	"github.com/gin-gonic/gin"
)

const (
	AvailableStatus  = "available"
	UnvailableStatus = "unavailable"
)

// @Tags HealthCheck
// @Summary Testar a conectividade da API
// @Description Testa a conectividade da API.
// @Produce json
// @Success 200 {object} viewmodels.HealthCheck
// @Router /health-check [get]
func HealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, viewmodels.HealthCheck{Status: AvailableStatus})
	}
}
