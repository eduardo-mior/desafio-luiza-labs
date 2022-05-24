package cep_controller

import (
	"desafio_luizalabs/domain"
	"desafio_luizalabs/domain/mocks"
	"desafio_luizalabs/viewmodels"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestBuscarCEPSucesso(t *testing.T) {

	serviceCEP := new(mocks.IServiceCEP)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Params = gin.Params{{Key: "cep", Value: "01001000"}}

	expectedCEP := domain.InformacoesCEP{UF: "SP", Cidade: "São Paulo", Bairro: "Sé", Rua: "Praça da Sé"}
	serviceCEP.On("BuscarCEPRecursivo", "01001000").Return(expectedCEP, nil)
	BuscarCEP(serviceCEP)(c)

	responseBody := viewmodels.InformacoesCEP{}
	err := json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, nil, err)

	expectedResponseCEP := viewmodels.InformacoesCEP{}
	expectedResponseCEP.FromDomainCEP(expectedCEP)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, expectedResponseCEP, responseBody)

}

func TestBuscarCEPErroCepInvalido1(t *testing.T) {

	serviceCEP := new(mocks.IServiceCEP)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Params = gin.Params{{Key: "cep", Value: "abcdefghj"}}

	expectedError := viewmodels.Error{Error: "CEP inválido"}

	BuscarCEP(serviceCEP)(c)

	responseBody := viewmodels.Error{}
	err := json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, nil, err)

	assert.Equal(t, 400, response.Code)
	assert.Equal(t, expectedError, responseBody)

}

func TestBuscarCEPErroCepInvalido2(t *testing.T) {

	serviceCEP := new(mocks.IServiceCEP)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Params = gin.Params{{Key: "cep", Value: "99150000"}}

	expectedError := viewmodels.Error{Error: "CEP inválido"}
	err := errors.New(expectedError.Error)

	serviceCEP.
		On("BuscarCEPRecursivo", "99150000").Return(domain.InformacoesCEP{}, err).
		On("IsErrCepInvalido", err).Return(true)

	BuscarCEP(serviceCEP)(c)

	responseBody := viewmodels.Error{}
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, nil, err)

	assert.Equal(t, 400, response.Code)
	assert.Equal(t, expectedError, responseBody)

}

func TestBuscarCEPErroInesperado(t *testing.T) {

	serviceCEP := new(mocks.IServiceCEP)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Params = gin.Params{{Key: "cep", Value: "99150000"}}

	expectedError := viewmodels.Error{Error: "Erro interno ao tentar buscar as informações do CEP"}
	err := errors.New(expectedError.Error)

	serviceCEP.
		On("BuscarCEPRecursivo", "99150000").Return(domain.InformacoesCEP{}, err).
		On("IsErrCepInvalido", err).Return(false)

	BuscarCEP(serviceCEP)(c)

	responseBody := viewmodels.Error{}
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, nil, err)

	assert.Equal(t, 500, response.Code)
	assert.Equal(t, expectedError, responseBody)

}
