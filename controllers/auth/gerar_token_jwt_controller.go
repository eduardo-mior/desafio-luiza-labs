package auth_controller

import (
	"desafio_luizalabs/domain"
	"desafio_luizalabs/viewmodels"

	"github.com/gin-gonic/gin"
)

const (
	erroAoGerarToken = "Erro interno ao tentar gerar o token JWT"
)

func GerarTokenJWT(authService domain.IServiceAuth) gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := authService.GenerateTokenJWT()
		if err != nil {
			c.AbortWithStatusJSON(500, viewmodels.Error{Error: erroAoGerarToken})
			return
		}

		c.JSON(200, viewmodels.Token{Token: token})
	}
}
