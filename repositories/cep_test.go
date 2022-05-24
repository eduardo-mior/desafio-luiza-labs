package repositories

import (
	"desafio_luizalabs/domain"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuscarCEPSucesso(t *testing.T) {

	repositoryCEP := RepositoryCEP{}

	cep, err := repositoryCEP.BuscarCEP("01001000")
	assert.Equal(t, nil, err)

	expectedCEP := domain.InformacoesCEP{UF: "SP", Cidade: "São Paulo", Bairro: "Sé", Rua: "Praça da Sé"}

	assert.Equal(t, expectedCEP, cep)

}

func TestBuscarCEPErro(t *testing.T) {

	repositoryCEP := RepositoryCEP{}

	_, err := repositoryCEP.BuscarCEP("00000000")
	assert.Equal(t, err, errors.New(errCepNaoEncontrado))

}

func TestIsErrCepNaoEncontrado(t *testing.T) {

	repositoryCEP := RepositoryCEP{}

	err := errors.New(errCepNaoEncontrado)
	isErrCepNaoEncontrado := repositoryCEP.IsErrCepNaoEncontrado(err)

	assert.Equal(t, isErrCepNaoEncontrado, true)

}
