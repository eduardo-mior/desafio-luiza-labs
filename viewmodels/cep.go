package viewmodels

import "desafio_luizalabs/domain"

type InformacoesCEP struct {
	UF     string `json:"uf"     example:"RS"`
	Cidade string `json:"cidade" example:"Marau"`
	Bairro string `json:"bairro" example:"Av. Julio Borella"`
	Rua    string `json:"rua"    example:"Rua costa silva"`
} // @name InformacoesCEP

func (informacoesCEP *InformacoesCEP) FromDomainCEP(domainCEP domain.InformacoesCEP) {
	informacoesCEP.UF = domainCEP.UF
	informacoesCEP.Cidade = domainCEP.Cidade
	informacoesCEP.Bairro = domainCEP.Bairro
	informacoesCEP.Rua = domainCEP.Rua
}
