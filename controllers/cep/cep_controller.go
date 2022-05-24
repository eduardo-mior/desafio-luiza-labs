package cep_controller

import (
	"desafio_luizalabs/domain"
	"desafio_luizalabs/util"
	"desafio_luizalabs/viewmodels"

	"github.com/gin-gonic/gin"
)

const (
	erroAoBuscarCEP = "Erro interno ao tentar buscar as informações do CEP"
	errCepInvalido  = "CEP inválido"
)

// @Tags CEP
// @Summary Buscar as informações de um CEP
// @Description Busca todas as informações de um cep retornando o estado (UF) o nome da cidade, o bairro e a rua de um CEP.
// @Produce json
// @Security ApiKeyAuth
// @Param cep path string true "CEP da localidade" minlength(8) maxlength(8)
// @Success 200 {object} viewmodels.InformacoesCEP
// @Failure 400,500 {object} viewmodels.Error
// @Failure 401 {object} viewmodels.AuthError
// @Router /buscar-cep/{cep} [get]
func BuscarCEP(serviceCEP domain.IServiceCEP) gin.HandlerFunc {
	return func(c *gin.Context) {

		cep := c.Param("cep")
		if len(cep) != 8 {
			c.AbortWithStatusJSON(400, viewmodels.Error{Error: errCepInvalido})
			return
		}

		cepInformacoes, err := serviceCEP.BuscarCEPRecursivo(cep)
		if err != nil {

			if serviceCEP.IsErrCepInvalido(err) {
				c.AbortWithStatusJSON(400, viewmodels.Error{Error: errCepInvalido})
			} else {
				util.GravarErroNoSentry(err, c)
				c.AbortWithStatusJSON(500, viewmodels.Error{Error: erroAoBuscarCEP})
			}

			return
		}

		cepResponse := viewmodels.InformacoesCEP{}
		cepResponse.FromDomainCEP(cepInformacoes)

		c.JSON(200, cepResponse)
	}
}
