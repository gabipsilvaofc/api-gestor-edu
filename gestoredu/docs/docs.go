// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/user": {
            "post": {
                "description": "Faz a criação de usuários",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Usuários"
                ],
                "summary": "Criar usuários",
                "parameters": [
                    {
                        "description": "dados do usuario",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Users"
                        }
                    },
                    "400": {
                        "description": "error"
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "description": "Faz a busca de usuários por ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Usuários"
                ],
                "summary": "Buscar usuários",
                "parameters": [
                    {
                        "description": "ID",
                        "name": "id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Users"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deleta usuário por ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Usuários"
                ],
                "summary": "Deletar usuários",
                "parameters": [
                    {
                        "description": "ID",
                        "name": "id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Users"
                        }
                    }
                }
            },
            "patch": {
                "description": "Atualiza dados do usuário por ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Usuários"
                ],
                "summary": "Atualizar usuários",
                "parameters": [
                    {
                        "description": "ID",
                        "name": "id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Users"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entities.UserRequest": {
            "type": "object",
            "required": [
                "cpf",
                "dataNascimento",
                "email",
                "endereco",
                "idade",
                "matricula",
                "nome",
                "sexo",
                "sobrenome",
                "telefone"
            ],
            "properties": {
                "cpf": {
                    "description": "User's CPF.\n@json\n@jsonTag cpf\n@jsonExample 000.000.000-00",
                    "type": "string",
                    "maxLength": 14,
                    "minLength": 11,
                    "example": "000.000.000-00"
                },
                "dataNascimento": {
                    "description": "User's dataNascimento (required, must be between 1 and 100).\n@json\n@jsonTag dataNascimento\n@jsonExample 30/03/2003",
                    "type": "string",
                    "maxLength": 10,
                    "minLength": 8,
                    "example": "30/03/2003"
                },
                "email": {
                    "description": "User's email (required and must be a valid email address).\nExample: user@example.com\n@json\n@jsonTag email\n@jsonExample user@example.com\n@binding required,email",
                    "type": "string",
                    "example": "test@test.com"
                },
                "endereco": {
                    "description": "User's endereco.\n@json\n@jsonTag endereco\n@jsonExample Rua Francisco Matos, 131\n@binding required,min=3,max=100",
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 3
                },
                "idade": {
                    "description": "User's idade (required, must be between 1 and 100).\n@json\n@jsonTag idade\n@jsonExample 30",
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1,
                    "example": "30"
                },
                "matricula": {
                    "description": "User's matricula.\n@json\n@jsonTag matricula\n@jsonExample 012345678",
                    "type": "string",
                    "maxLength": 10,
                    "minLength": 8,
                    "example": "012345678"
                },
                "nome": {
                    "description": "User's name (required, minimum of 4 characters, maximum of 100 characters).\n@json\n@jsonTag name\n@binding required,min=4,max=100",
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 4
                },
                "sexo": {
                    "description": "User's\n@json\n@jsonTag sexo\n@jsonExample F\n@binding required,sexo",
                    "type": "string",
                    "example": "F"
                },
                "sobrenome": {
                    "description": "User's sobrenome (required, minimum of 4 characters, maximum of 100 characters).\n@json\n@jsonTag sobrenome\n@binding required,min=4,max=100",
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 4
                },
                "telefone": {
                    "description": "User's telefone.\n@json\n@jsonTag telefone\n@jsonExample 48988881111",
                    "type": "string",
                    "maxLength": 11,
                    "minLength": 8,
                    "example": "48988881111"
                }
            }
        },
        "entities.Users": {
            "type": "object",
            "properties": {
                "cpf": {
                    "type": "string"
                },
                "dataNascimento": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "endereco": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "idade": {
                    "type": "string"
                },
                "matricula": {
                    "type": "string"
                },
                "nome": {
                    "type": "string"
                },
                "sexo": {
                    "type": "string"
                },
                "sobrenome": {
                    "type": "string"
                },
                "telefone": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "API GestorEdu",
	Description:      "Realiza CRUD de usuarios",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
