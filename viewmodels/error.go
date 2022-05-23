package viewmodels

type Error struct {
	Error string `json:"error" example:"CEP inválido"`
}

type AuthError struct {
	Error string `json:"authError" example:"Token de autenticação inválido"`
}
