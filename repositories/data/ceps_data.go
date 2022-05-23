package data

import "desafio_luizalabs/domain"

var CEPS = map[string]domain.InformacoesCEP{
	"01001000": {UF: "SP", Cidade: "São Paulo", Bairro: "Sé", Rua: "Praça da Sé"},
	"01001001": {UF: "SP", Cidade: "São Paulo", Bairro: "Sé", Rua: "Praça da Sé"},
	"01153000": {UF: "SP", Cidade: "São Paulo", Bairro: "Barra Funda", Rua: "Rua Vitorino Carmilo"},
	"03223050": {UF: "SP", Cidade: "São Paulo", Bairro: "Jardim Independência", Rua: "Rua Dario Meira"},
	"03255000": {UF: "SP", Cidade: "São Paulo", Bairro: "Vila Tolstoi", Rua: "Rua José Antônio Fontes"},
	"03513050": {UF: "SP", Cidade: "São Paulo", Bairro: "Vila Matilde", Rua: "Rua Antônio Juvenal"},
	"03683010": {UF: "SP", Cidade: "São Paulo", Bairro: "Vila União (Zona Leste)", Rua: "Rua Antônio Olímpio"},
	"04295001": {UF: "SP", Cidade: "São Paulo", Bairro: "Vila Vera", Rua: "Avenida Coronel José Pires de Andrade"},
	"05425070": {UF: "SP", Cidade: "São Paulo", Bairro: "Pinheiros", Rua: "Avenida Doutora Ruth Cardoso"},
	"05434020": {UF: "SP", Cidade: "São Paulo", Bairro: "Vila Madalena", Rua: "Rua Morás"},
	"08090284": {UF: "SP", Cidade: "São Paulo", Bairro: "Jardim Helena", Rua: "Rua 03 de Outubro"},
	"08410500": {UF: "SP", Cidade: "São Paulo", Bairro: "Vila Santa Cruz (Zona Leste)", Rua: "Rua Doutor Pedro Batista"},
	"09925000": {UF: "SP", Cidade: "Diadema", Bairro: "Campanário", Rua: "Rodovia dos Imigrantes"},
	"11441710": {UF: "SP", Cidade: "Guarujá", Bairro: "Enseada", Rua: "Rua Projetada Nove"},
	"12209530": {UF: "SP", Cidade: "São José dos Campos", Bairro: "Centro", Rua: "Rua José de Alencar"},
	"13327220": {UF: "SP", Cidade: "Salto", Bairro: "Jardim Saltense", Rua: "Rua Acácio Rodrigues de Moraes"},
	"13401130": {UF: "SP", Cidade: "Piracicaba", Bairro: "Paulista", Rua: "Rua Sud Mennucci"},
	"17022113": {UF: "SP", Cidade: "Bauru", Bairro: "Vila São Paulo", Rua: "Rua Gaudêncio Piola"},
	"20020030": {UF: "RJ", Cidade: "Rio de Janeiro", Bairro: "Centro", Rua: "Praça Academia"},
	"20760245": {UF: "RJ", Cidade: "Rio de Janeiro", Bairro: "Inhaúma", Rua: "Rua José dos Reis"},
	"22610001": {UF: "RJ", Cidade: "Rio de Janeiro", Bairro: "São Conrado", Rua: "Estrada da Gávea"},
	"23065000": {UF: "RJ", Cidade: "Rio de Janeiro", Bairro: "Paciência", Rua: "Estrada Manguariba"},
	"23587608": {UF: "RJ", Cidade: "Rio de Janeiro", Bairro: "Paciência", Rua: "Rua Graça e Paz"},
	"24220091": {UF: "RJ", Cidade: "Niterói", Bairro: "Icaraí", Rua: "Rua Domingues de Sá"},
	"27913300": {UF: "RJ", Cidade: "Macaé", Bairro: "Imbetiba", Rua: "Avenida Agenor Caldas"},
	"29010100": {UF: "ES", Cidade: "Vitória", Bairro: "Centro", Rua: "Rua Dionísio Rosendo"},
	"29216170": {UF: "ES", Cidade: "Guarapari", Bairro: "Praia do Morro", Rua: "Rua Saint Tropez"},
	"30150313": {UF: "MG", Cidade: "Belo Horizonte", Bairro: "Funcionários", Rua: "Rua Ceará"},
	"30315600": {UF: "MG", Cidade: "Belo Horizonte", Bairro: "Sion", Rua: "Rua Grão Mogol"},
	"30421052": {UF: "MG", Cidade: "Belo Horizonte", Bairro: "Nova Suíssa", Rua: "Rua Genebra"},
	"30530000": {UF: "MG", Cidade: "Belo Horizonte", Bairro: "João Pinheiro", Rua: "Avenida Vereador Cícero Ildefonso"},
	"31814650": {UF: "MG", Cidade: "Belo Horizonte", Bairro: "Aarão Reis", Rua: "Rua Irmãos Rodrigues"},
	"31970420": {UF: "MG", Cidade: "Belo Horizonte", Bairro: "Jardim Vitória", Rua: "Rua Maria Jacinta"},
	"33010620": {UF: "MG", Cidade: "Santa Luzia", Bairro: "Bela Vista", Rua: "Rua Palmor Teixeira Viana"},
	"33045370": {UF: "MG", Cidade: "Santa Luzia", Bairro: "Frimisa", Rua: "Rua R"},
	"33822720": {UF: "MG", Cidade: "Ribeirão das Neves", Bairro: "Liberdade", Rua: "Avenida B"},
	"35500538": {UF: "MG", Cidade: "Divinópolis", Bairro: "Santa Rosa", Rua: "Rua Moscou"},
	"35700292": {UF: "MG", Cidade: "Sete Lagoas", Bairro: "Canaã", Rua: "Rua Natal"},
	"36774406": {UF: "MG", Cidade: "Cataguases", Bairro: "Sol Nascente", Rua: "Rua Álvaro Franca"},
	"38400710": {UF: "MG", Cidade: "Uberlândia", Bairro: "Brasil", Rua: "Avenida Afonso Pena"},
	"38703240": {UF: "MG", Cidade: "Patos de Minas", Bairro: "Bela Vista", Rua: "Rua Arnaldo Luiz de Oliveira"},
	"39803401": {UF: "MG", Cidade: "Teófilo Otoni", Bairro: "Jardim São Paulo", Rua: "Rua São Caetano"},
	"44028502": {UF: "BA", Cidade: "Feira de Santana", Bairro: "Gabriela", Rua: "Rua Professora Maria da Glória"},
	"44054394": {UF: "BA", Cidade: "Feira de Santana", Bairro: "Parque Ipê", Rua: "Rua Tijuca"},
	"45656220": {UF: "BA", Cidade: "Ilhéus", Bairro: "Nelson Costa", Rua: "Rua Bem-Me-Quer"},
	"53401000": {UF: "PE", Cidade: "Paulista", Bairro: "Centro", Rua: "Rua Campo Grande"},
	"54505470": {UF: "PE", Cidade: "Cabo de Santo Agostinho", Bairro: "Centro", Rua: "Rua Vigário João Batista"},
	"57306370": {UF: "AL", Cidade: "Arapiraca", Bairro: "Cavaco", Rua: "Rua Manoel Ferreira de Brito"},
	"58832970": {UF: "PB", Cidade: "Mato Grosso", Bairro: "Centro", Rua: "Rua João Serafim de Lima 96"},
	"71225520": {UF: "DF", Cidade: "Brasília", Bairro: "Zona Industrial (Guará)", Rua: "Trecho STRC Trecho 2"},
	"72430204": {UF: "DF", Cidade: "Brasília", Bairro: "Setor Norte (Gama)", Rua: "Quadra 2 Conjunto D"},
	"74910095": {UF: "GO", Cidade: "Aparecida de Goiânia", Bairro: "Parque Real de Goiânia", Rua: "Avenida São Paulo"},
	"78557165": {UF: "MT", Cidade: "Sinop", Bairro: "Setor Industrial", Rua: "Rua João Pedro Moreira de Carvalho"},
	"78559889": {UF: "MT", Cidade: "Sinop", Bairro: "Residencial sabrina III", Rua: "Rua Projetada 2"},
	"78559899": {UF: "MT", Cidade: "Sinop", Bairro: "Área Rural de Sinop", Rua: "Área Rural"},
	"79106400": {UF: "MS", Cidade: "Campo Grande", Bairro: "Vila Popular", Rua: "Rua Domingos Aparecido Bissoli"},
	"79116121": {UF: "MS", Cidade: "Campo Grande", Bairro: "Parque Residencial Azaléia", Rua: "Rua Bauxi"},
	"79630090": {UF: "MS", Cidade: "Três Lagoas", Bairro: "Jardim das Oliveiras", Rua: "Rua Protázio Garcia Leal"},
	"80410201": {UF: "PR", Cidade: "Curitiba", Bairro: "Centro", Rua: "Rua Visconde de Nacar"},
	"83823138": {UF: "PR", Cidade: "Fazenda Rio Grande", Bairro: "Nações", Rua: "Avenida Paraguai"},
	"85501240": {UF: "PR", Cidade: "Pato Branco", Bairro: "Centro", Rua: "Rua Itacolomi"},
	"85806077": {UF: "PR", Cidade: "Cascavel", Bairro: "FAG", Rua: "Rua João Kuloski"},
	"85813390": {UF: "PR", Cidade: "Cascavel", Bairro: "Country", Rua: "Rua Acre"},
	"87112257": {UF: "PR", Cidade: "Sarandi", Bairro: "Parque Residencial Bela Vista", Rua: "Avenida Amambai"},
	"89226500": {UF: "SC", Cidade: "Joinville", Bairro: "Jardim Paraíso", Rua: "Estrada Timbé"},
	"95700010": {UF: "RS", Cidade: "Bento Gonçalves", Bairro: "Centro", Rua: "Rua Marechal Deodoro"},
	"99010000": {UF: "RS", Cidade: "Passo Fundo", Bairro: "Centro", Rua: "Avenida Brasil"},
	"99010030": {UF: "RS", Cidade: "Passo Fundo", Bairro: "Centro", Rua: "Rua Morom"},
	"99030010": {UF: "RS", Cidade: "Passo Fundo", Bairro: "Boqueirão", Rua: "Rua Tonico Silva"},
}
