package middlewares

import (
	"desafio_luizalabs/viewmodels"

	"github.com/gin-gonic/gin"
)

const (
	Erro404RecursoNaoEncontrado = "Recurso não encontrado"
)

// NotFound é o Middleware responsável por retornar um JSON com erro 404, nós criamos ele para substituir o "404 page not found" padrão do Gin.
func NotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.AbortWithStatusJSON(404, viewmodels.Error{Error: Erro404RecursoNaoEncontrado})
	}
}
