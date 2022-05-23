package middlewares

import (
	"desafio_luizalabs/viewmodels"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const (
	erroTokenNaoInformado = "Token de autenticação não informado"
	erroTokenInvalido     = "Token de autenticação inválido"
)

// AuthJWT é o Middleware responsável por validar a autenticidade de uma requisição através do seu token JWT.
func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		headerAuthorization := c.GetHeader("Authorization")
		if headerAuthorization == "" {
			c.AbortWithStatusJSON(401, viewmodels.AuthError{Error: erroTokenNaoInformado})
			return
		}

		bearerToken := strings.Split(headerAuthorization, "Bearer ")
		if len(bearerToken) != 2 {
			c.AbortWithStatusJSON(401, viewmodels.AuthError{Error: erroTokenInvalido})
			return
		}

		tokenString := bearerToken[1]

		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(401, viewmodels.AuthError{Error: erroTokenInvalido})
			return
		}

		c.Next()
	}
}
