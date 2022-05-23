package repositories

import (
	"desafio_luizalabs/domain"
	"desafio_luizalabs/repositories/data"
	"errors"
)

type RepositoryCEP struct{}

const (
	errCepNaoEncontrado = "CEP não encontrado"
)

func (RepositoryCEP) BuscarCEP(cep string) (domain.InformacoesCEP, error) {

	informacoesCEP, existe := data.CEPS[cep]
	if existe {
		return informacoesCEP, nil
	}

	return domain.InformacoesCEP{}, errors.New(errCepNaoEncontrado)
}

func (RepositoryCEP) IsErrCepNaoEncontrado(err error) bool {
	return err != nil && err.Error() == errCepNaoEncontrado
}
