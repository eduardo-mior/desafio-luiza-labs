# Desafio LuizaLabs
Implementação do Desafio Técnico da LuizaLabs.

## Linguagem
A linguagem escolhida para desenvolver o desafio foi a linguagem [GO](https://go.dev/) pelo fato de que GO é uma linguagem de programação simples, onde os códigos podem ser lidos e escritos com facilidade, e que além disso possui uma alta performance e escalabilidade.

## Arquitetura
A arquitetura implementada no projeto segue os conceitos da CleanArchitecture e os princípios SOLID, com baixo acoplamento, separação de interesses, injeção de dependências, mock de dados e testes unitários.

Todo o trafego de dados ocorre no padrão `application/json`.

Para questões de segurança e autorização foi utilizado o padrão de tokens JWT.

## Ferramentas
Para o desenvolvimento e implantação do projeto foram usadas as seguintes ferramentas:
- [GIN](https://github.com/gin-gonic/gin): Framework WEB para tratamento das requisições HTTP.
- [JWT](github.com/golang-jwt/jwt): Geração e validação dos tokens JWT.
- [Swaggo](github.com/swaggo/gin-swagger): Geração automática e disponibilização da documentação no padrão OpenAPI com Swagger.
- [Testify](https://github.com/stretchr/testify): Conjunto de pacotes e funções utilitárias que ajudam no desenvolvimento dos testes unitários.
- [Mockery](https://github.com/vektra/mockery): Gerador de Mocks para facilitar o desenvolvimento dos testes unitários.
- [Postman](https://www.postman.com/): Ferramenda usada para criar os testes de ponta a ponta (e2e).