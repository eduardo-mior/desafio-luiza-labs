# Desafio LuizaLabs
Implementação do Desafio Técnico da LuizaLabs.

#
# 1 - Serviço de busca de CEP

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

# 
# 2 - Questão teórica
Questão: `Quando você digita a URL de um site (http://www.netshoes.com.br) no browser e pressiona enter, explique da forma que preferir, o que ocorre nesse processo do protocolo HTTP entre o Client e o Server?`

Resposta: Ao acessar o site o http://www.netshoes.com.br o browser faz uma requisição do tipo GET para o servidor da netshoes que retorna uma resposta com StatusCode 307 (Internal Redirect) indicando ao Browser (Client) que é necessário fazer um redirecionamento, além do StatusCode 307 o servidor também retornou um cabeçalho chamado `Location`, este cabeçalho contém a URL de destinado para onde o browser deve redirecionar o usuário. 

Neste caso da netshoes em específico esta acontecendo um redicionamento do protocolo http para o https, este é um redirecionamento muito comum, ele é feito por questões de segurança pelo fato de que o procolo http não ser considerado seguro.

Após o redirecionamento ser feito o Browser começa a fazer uma nova requisição para a nova URL e nesse caso o servidor retorna um StatusCode 200 (Sucesso) e também retora uma texto/html como resposta no corpo do retorno. 

Após o Browser ter recebido com sucesso todo conteúdo da resposta ele começa a renderizar este conteúdo na tela e a partir desse momento o Browser também começa a fazer uma série de novas requisições para baixar os JavaScripts, CSSs, Fontes, Imagens e outros Assets do site até que ele seja renderizado por completo.

De forma resumida o Client que é o Browser faz requisições para o Servidor que responde com conteúdos de texto e mídia. 