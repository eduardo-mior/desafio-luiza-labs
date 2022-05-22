#####
# Build
#####
# Baixando a imagem do GO para realizar o Build do projeto
FROM golang:1.18.2-alpine AS build_base

# Definino o direitório de trabalho, a partir de agora todos os comandos serão executados a partir desse path
WORKDIR /app

# Copiando os fontes do projeto para a pasta o diretório de trabalho atual
COPY . ./

# Baixando dependências do projeto
RUN go get

# Compilando o projeto
RUN go build -o /output/desafio_luizalabs

#####
# Deploy
#####
# Baixando a imagem do AlpineLinux para executar o projeto
FROM alpine:3.15

# Copiando o executavel gerado no estagio de Build para poder executado
COPY --from=build_base /output/desafio_luizalabs /desafio_luizalabs

# Abrindo a porta 8080 padrão 
EXPOSE 8080

# Inicializando o container
CMD [ "/desafio_luizalabs" ]
