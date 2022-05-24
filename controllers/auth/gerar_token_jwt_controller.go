package auth_controller

import (
	"desafio_luizalabs/domain"
	"desafio_luizalabs/viewmodels"

	"github.com/gin-gonic/gin"
)

const (
	erroAoGerarToken = "Erro interno ao tentar gerar o token JWT"
)

// @Tags Auth
// @Summary Gerar um token JWT de testes
// @Description Gera um token JWT de testes válido por 1 dia.
// @Produce json
// @Success 200 {object} viewmodels.Token
// @Failure 400,500 {object} viewmodels.Error
// @Router /gerar-token-jwt [get]
func GerarTokenJWT(serviceAuth domain.IServiceAuth) gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := serviceAuth.GenerateTokenJWT()
		if err != nil {
			c.AbortWithStatusJSON(500, viewmodels.Error{Error: erroAoGerarToken})
			return
		}

		c.JSON(200, viewmodels.Token{Token: token})
	}
}
