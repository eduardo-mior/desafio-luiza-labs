package viewmodels

type Error struct {
	Error string `json:"error" example:"Ocorreu um erro ao tentar realizar a operação"`
} // @name Error

type AuthError struct {
	Error string `json:"authError" example:"Token de autenticação inválido"`
} // @name AuthError
