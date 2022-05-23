package middlewares

import (
	"desafio_luizalabs/viewmodels"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

const (
	erro500Interno = "Erro interno ao tentar realizar a operação"
)

// PanicRecovery é o Middleware responsável por recuperar uma requisição caso aconteça uma falha cadastrófica no Controller.
func PanicRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {

				debug.PrintStack()

				c.JSON(500, viewmodels.Error{Error: erro500Interno})

			}
		}()
		c.Next()
	}
}
