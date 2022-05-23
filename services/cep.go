package services

import (
	"desafio_luizalabs/domain"
	"errors"
	"strings"
)

type ServiceCEP struct {
	domain.IRepositoryCEP
}

const (
	errCepInvalido = "CEP inválido"
)

func (service ServiceCEP) BuscarCEPRecursivo(cep string) (domain.InformacoesCEP, error) {

	for i := 1; i < 8; i++ {

		cepInformacoes, err := service.BuscarCEP(cep)

		if err != nil {

			if service.IsErrCepNaoEncontrado(err) {
				cep = cep[:8-i] + strings.Repeat("0", i)
			} else {
				return domain.InformacoesCEP{}, err
			}

		} else {
			return cepInformacoes, nil
		}

	}

	return domain.InformacoesCEP{}, errors.New(errCepInvalido)
}

func (ServiceCEP) IsErrCepInvalido(err error) bool {
	return err != nil && err.Error() == errCepInvalido
}
