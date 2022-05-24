// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/buscar-cep/{cep}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Busca todas as informações de um cep retornando o estado (UF) o nome da cidade, o bairro e a rua de um CEP.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CEP"
                ],
                "summary": "Buscar as informações de um CEP",
                "parameters": [
                    {
                        "maxLength": 8,
                        "minLength": 8,
                        "type": "string",
                        "description": "CEP da localidade",
                        "name": "cep",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/InformacoesCEP"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/AuthError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/Error"
                        }
                    }
                }
            }
        },
        "/gerar-token-jwt": {
            "get": {
                "description": "Gera um token JWT de testes válido por 1 dia.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Gerar um token JWT de testes",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Token"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/Error"
                        }
                    }
                }
            }
        },
        "/health-check": {
            "get": {
                "description": "Testa a conectividade da API.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HealthCheck"
                ],
                "summary": "Testar a conectividade da API",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/HealthCheck"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "AuthError": {
            "type": "object",
            "properties": {
                "authError": {
                    "type": "string",
                    "example": "Token de autenticação inválido"
                }
            }
        },
        "Error": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Ocorreu um erro ao tentar realizar a operação"
                }
            }
        },
        "HealthCheck": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "example": "available"
                }
            }
        },
        "InformacoesCEP": {
            "type": "object",
            "properties": {
                "bairro": {
                    "type": "string",
                    "example": "Av. Julio Borella"
                },
                "cidade": {
                    "type": "string",
                    "example": "Marau"
                },
                "rua": {
                    "type": "string",
                    "example": "Rua costa silva"
                },
                "uf": {
                    "type": "string",
                    "example": "RS"
                }
            }
        },
        "Token": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UifQ.b-bQWDLukCsP-kz255shgiDNoPjycxSizc1cIljoEic"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}