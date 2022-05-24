package services

import (
	"desafio_luizalabs/domain"
	"desafio_luizalabs/domain/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuscarCEPSucesso(t *testing.T) {

	expectedCEP := domain.InformacoesCEP{UF: "SP", Cidade: "São Paulo", Bairro: "Sé", Rua: "Praça da Sé"}

	repositoryCEP := new(mocks.IRepositoryCEP)
	repositoryCEP.
		On("BuscarCEP", "01001999").Return(domain.InformacoesCEP{}, errors.New("CEP não encontrado")).
		On("IsErrCepNaoEncontrado", errors.New("CEP não encontrado")).Return(true).
		On("BuscarCEP", "01001990").Return(domain.InformacoesCEP{}, errors.New("CEP não encontrado")).
		On("IsErrCepNaoEncontrado", errors.New("CEP não encontrado")).Return(true).
		On("BuscarCEP", "01001900").Return(domain.InformacoesCEP{}, errors.New("CEP não encontrado")).
		On("IsErrCepNaoEncontrado", errors.New("CEP não encontrado")).Return(true).
		On("BuscarCEP", "01001000").Return(expectedCEP, nil)

	serviceCEP := ServiceCEP{repositoryCEP}
	cep, err := serviceCEP.BuscarCEPRecursivo("01001999")
	assert.Equal(t, nil, err)

	assert.Equal(t, expectedCEP, cep)

}

func TestBuscarCEPInvalido(t *testing.T) {

	repositoryCEP := new(mocks.IRepositoryCEP)
	repositoryCEP.
		On("BuscarCEP", "00000000").Return(domain.InformacoesCEP{}, errors.New("CEP não encontrado")).
		On("IsErrCepNaoEncontrado", errors.New("CEP não encontrado")).Return(true)

	serviceCEP := ServiceCEP{repositoryCEP}
	_, err := serviceCEP.BuscarCEPRecursivo("00000000")
	assert.Equal(t, err, errors.New(errCepInvalido))

}

func TestBuscarCEPErroDesconhecido(t *testing.T) {

	repositoryCEP := new(mocks.IRepositoryCEP)
	repositoryCEP.
		On("BuscarCEP", "00000000").Return(domain.InformacoesCEP{}, errors.New("Erro interno desconhecido...")).
		On("IsErrCepNaoEncontrado", errors.New("Erro interno desconhecido...")).Return(false)

	serviceCEP := ServiceCEP{repositoryCEP}
	_, err := serviceCEP.BuscarCEPRecursivo("00000000")
	assert.Equal(t, err, errors.New("Erro interno desconhecido..."))

}

func TestIsErrCepInvalido(t *testing.T) {
	serviceCEP := ServiceCEP{}

	err := errors.New(errCepInvalido)
	isErrCepInvalido := serviceCEP.IsErrCepInvalido(err)

	assert.Equal(t, isErrCepInvalido, true)
}
