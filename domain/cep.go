package domain

type IServiceCEP interface {
	BuscarCEPRecursivo(cep string) (InformacoesCEP, error)
	IsErrCepInvalido(error) bool
}

type IRepositoryCEP interface {
	BuscarCEP(cep string) (InformacoesCEP, error)
	IsErrCepNaoEncontrado(error) bool
}

type InformacoesCEP struct {
	UF     string
	Cidade string
	Bairro string
	Rua    string
}
