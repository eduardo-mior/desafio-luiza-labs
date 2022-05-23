package healthcheck_controller

import (
	"desafio_luizalabs/viewmodels"

	"github.com/gin-gonic/gin"
)

const (
	AvailableStatus  = "available"
	UnvailableStatus = "unavailable"
)

func HealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, viewmodels.HealthCheck{Status: AvailableStatus})
	}
}
