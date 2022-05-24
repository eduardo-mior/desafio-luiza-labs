package cep_controller

import (
	"desafio_luizalabs/domain"
	"desafio_luizalabs/viewmodels"

	"github.com/gin-gonic/gin"
)

const (
	erroAoBuscarCEP = "Erro interno ao tentar buscar as informações do CEP"
)

func BuscarCEP(serviceCEP domain.IServiceCEP) gin.HandlerFunc {
	return func(c *gin.Context) {

		cep := c.Param("cep")

		cepInformacoes, err := serviceCEP.BuscarCEPRecursivo(cep)
		if err != nil {

			if serviceCEP.IsErrCepInvalido(err) {
				c.AbortWithStatusJSON(400, viewmodels.Error{Error: err.Error()})
			} else {
				c.AbortWithStatusJSON(500, viewmodels.Error{Error: erroAoBuscarCEP})
			}

			return
		}

		cepResponse := viewmodels.InformacoesCEP{}
		cepResponse.FromDomainCEP(cepInformacoes)

		c.JSON(200, cepResponse)
	}
}
