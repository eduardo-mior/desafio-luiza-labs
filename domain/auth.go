package domain

type IServiceAuth interface {
	GenerateTokenJWT() (string, error)
}
