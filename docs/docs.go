// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Unknown"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/users": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Gets a list of all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/modelcrud.APIUser"
                            }
                        }
                    }
                }
            }
        },
        "/auth/loginjwt": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Exchanges a JWT from a configurable source for a signed JWT",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/auth/loginup": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Exchanges a username and password for a signed JWT",
                "parameters": [
                    {
                        "description": "Login parameters",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Registers a new account using username and password as the authentication scheme",
                "parameters": [
                    {
                        "description": "Register parameters",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/crud/accessLogs": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Gets a list for all entities of the AccessLog type",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/modelcrud.APIAccessLog"
                        }
                    }
                }
            }
        },
        "/crud/accessLogs/:id": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Gets a single AccessLog entity by their id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AccessLog id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/modelcrud.APIGetAccessLogsResponse"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Updates a single AccessLog entity based on their id",
                "parameters": [
                    {
                        "description": "Update parameters",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/modelcrud.APIUpdateAccessLogRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "AccessLog id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/modelcrud.APIAccessLog"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "Soft deletes a single AccessLog entity based on their id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AccessLog id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/modelcrud.APIAccessLog"
                        }
                    }
                }
            }
        },
        "/crud/tokenIssuances": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Gets a list for all entities of the TokenIssuance type",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/modelcrud.APITokenIssuance"
                        }
                    }
                }
            }
        },
        "/crud/tokenIssuances/:id": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Gets a single TokenIssuance entity by their id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "TokenIssuance id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/modelcrud.APIGetTokenIssuancesResponse"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Updates a single TokenIssuance entity based on their id",
                "parameters": [
                    {
                        "description": "Update parameters",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/modelcrud.APIUpdateTokenIssuanceRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "TokenIssuance id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/modelcrud.APITokenIssuance"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "Soft deletes a single TokenIssuance entity based on their id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "TokenIssuance id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/modelcrud.APITokenIssuance"
                        }
                    }
                }
            }
        },
        "/crud/users": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Gets a list for all entities of the User type",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Include deleted users in the results",
                        "name": "withDeleted",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/modelcrud.APIUser"
                        }
                    }
                }
            }
        },
        "/crud/users/:id": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Gets a single User entity by their id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/modelcrud.APIGetUsersResponse"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Updates a single User entity based on their id",
                "parameters": [
                    {
                        "description": "Update parameters",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/modelcrud.APIUpdateUserRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "User id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/modelcrud.APIUser"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "Soft deletes a single User entity based on their id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Hard delete user",
                        "name": "hardDelete",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/modelcrud.APIUser"
                        }
                    }
                }
            }
        },
        "/healthz": {
            "get": {
                "description": "returns 200 when service is running",
                "produces": [
                    "application/json"
                ],
                "summary": "Health check handler",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/users/me": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Gets information on the current user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/modelcrud.APIUser"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.LoginRequest": {
            "type": "object",
            "required": [
                "login",
                "password"
            ],
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "main.RegisterRequest": {
            "type": "object",
            "required": [
                "email",
                "login",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "login": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "modelcrud.APIAccessLog": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "ip_address": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "processing_duration": {
                    "type": "integer"
                },
                "request_body": {
                    "type": "string"
                },
                "request_headers": {
                    "type": "object",
                    "additionalProperties": true
                },
                "request_method": {
                    "type": "string"
                },
                "response_body": {
                    "type": "object",
                    "additionalProperties": true
                },
                "response_code": {
                    "type": "integer"
                },
                "response_headers": {
                    "type": "object",
                    "additionalProperties": true
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "modelcrud.APIGetAccessLogsResponse": {
            "type": "object",
            "properties": {
                "access_logs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/modelcrud.APIAccessLog"
                    }
                },
                "next_offset": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "modelcrud.APIGetTokenIssuancesResponse": {
            "type": "object",
            "properties": {
                "next_offset": {
                    "type": "integer"
                },
                "token_issuances": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/modelcrud.APITokenIssuance"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "modelcrud.APIGetUsersResponse": {
            "type": "object",
            "properties": {
                "next_offset": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/modelcrud.APIUser"
                    }
                }
            }
        },
        "modelcrud.APITokenIssuance": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "ip_address": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "modelcrud.APIUpdateAccessLogRequest": {
            "type": "object",
            "properties": {
                "ip_address": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "processing_duration": {
                    "type": "integer"
                },
                "request_body": {
                    "type": "string"
                },
                "request_headers": {
                    "type": "object",
                    "additionalProperties": true
                },
                "request_method": {
                    "type": "string"
                },
                "response_body": {
                    "type": "object",
                    "additionalProperties": true
                },
                "response_code": {
                    "type": "integer"
                },
                "response_headers": {
                    "type": "object",
                    "additionalProperties": true
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "modelcrud.APIUpdateTokenIssuanceRequest": {
            "type": "object",
            "properties": {
                "ip_address": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "modelcrud.APIUpdateUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "login": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "sub": {
                    "type": "string"
                }
            }
        },
        "modelcrud.APIUser": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "login": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "sub": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.1",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Go Vue Template",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
