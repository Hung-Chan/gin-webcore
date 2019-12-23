// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-12-23 07:22:40.9792823 +0000 UTC m=+0.140413501

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin-accesses/": {
            "get": {
                "description": "GET Admin Access List",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AdminAccesses"
                ],
                "summary": "Admin Access List",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "SortColumn",
                        "name": "sortColumn",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "SortDirection",
                        "name": "sortDirection",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Enable",
                        "name": "enable",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    }
                }
            },
            "post": {
                "description": "GET Admin Access Create",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AdminAccesses"
                ],
                "summary": "Admin Access Create",
                "parameters": [
                    {
                        "description": "Admin Access Create",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/adminaccesses.AdminAccessModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    }
                }
            }
        },
        "/admin-accesses/view/{id}": {
            "get": {
                "description": "GET Admin Access View",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AdminAccesses"
                ],
                "summary": "Admin Access View",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Admin Access ID",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    }
                }
            }
        },
        "/admin-accesses/{id}": {
            "delete": {
                "description": "DELETE Admin Access Delete",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AdminAccesses"
                ],
                "summary": "Admin Access Delete",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Admin Access ID",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    }
                }
            },
            "patch": {
                "description": "PATCH Admin Access Update",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AdminAccesses"
                ],
                "summary": "Admin Access Update",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Admin Access ID",
                        "name": "id",
                        "in": "path"
                    },
                    {
                        "description": "Admin Access Update",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/adminaccesses.AdminAccessModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    }
                }
            }
        },
        "/admin-levels/": {
            "get": {
                "description": "GET Admin Levels List",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AdminLevels"
                ],
                "summary": "Admin Levels List",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "SortColumn",
                        "name": "sortColumn",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "SortDirection",
                        "name": "sortDirection",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Enable",
                        "name": "enable",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    }
                }
            },
            "post": {
                "description": "GET Admin Level Create",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AdminLevels"
                ],
                "summary": "Admin Level Create",
                "parameters": [
                    {
                        "description": "Admin Level Create",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/adminlevels.AdminLevelModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    }
                }
            }
        },
        "/admin-levels/{id}": {
            "get": {
                "description": "GET Admin Level View",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AdminLevels"
                ],
                "summary": "Admin Level View",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Admin Level ID",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    }
                }
            },
            "delete": {
                "description": "GET Admin Level Delete",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AdminLevels"
                ],
                "summary": "Admin Level Delete",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Admin Level ID",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    }
                }
            },
            "patch": {
                "description": "GET Admin Level Update",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AdminLevels"
                ],
                "summary": "Admin Level Update",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Admin Level ID",
                        "name": "id",
                        "in": "path"
                    },
                    {
                        "description": "Admin Level Update",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/adminlevels.AdminLevelModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    }
                }
            }
        },
        "/auth/info": {
            "get": {
                "description": "Get Admin Info",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Admin Info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Admin Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Admin Login",
                "parameters": [
                    {
                        "description": "login",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/auth.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        },
                        "headers": {
                            "Token": {
                                "type": "string",
                                "description": "qwerty"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    }
                }
            }
        },
        "/auth/sidebarMenu": {
            "get": {
                "description": "Get SidebarMenu",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Admin SidebarMenu",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    }
                }
            }
        },
        "/ip-whitelistings/": {
            "get": {
                "description": "GET IP Whitelisting List",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "IPWhitelistings"
                ],
                "summary": "IP Whitelisting List",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "SortColumn",
                        "name": "sortColumn",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "SortDirection",
                        "name": "sortDirection",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Enable",
                        "name": "enable",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    }
                }
            },
            "post": {
                "description": "GET IP Whitelisting Create",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "IPWhitelistings"
                ],
                "summary": "IP Whitelisting Create",
                "parameters": [
                    {
                        "description": "IP Whitelisting Create",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/ipwhitelistings.IPWhitelistingModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    }
                }
            }
        },
        "/ip-whitelistings/view/{id}": {
            "get": {
                "description": "GET IP Whitelisting View",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "IPWhitelistings"
                ],
                "summary": "IP Whitelisting View",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "IP Whitelisting ID",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    }
                }
            }
        },
        "/ip-whitelistings/{id}": {
            "delete": {
                "description": "DELETE IP Whitelisting Delete",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "IPWhitelistings"
                ],
                "summary": "IP Whitelisting Delete",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "IP Whitelisting ID",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    }
                }
            },
            "patch": {
                "description": "PATCH IP Whitelisting Update",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "IPWhitelistings"
                ],
                "summary": "IP Whitelisting Update",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "IP Whitelisting ID",
                        "name": "id",
                        "in": "path"
                    },
                    {
                        "description": "IP Whitelisting Update",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/ipwhitelistings.IPWhitelistingModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "adminaccesses.AdminAccessModel": {
            "type": "object",
            "required": [
                "code",
                "name"
            ],
            "properties": {
                "code": {
                    "type": "string",
                    "example": "test"
                },
                "enable": {
                    "type": "integer"
                },
                "name": {
                    "type": "string",
                    "example": "test"
                }
            }
        },
        "adminlevels.AdminLevelModel": {
            "type": "object",
            "required": [
                "level",
                "name"
            ],
            "properties": {
                "enable": {
                    "type": "integer"
                },
                "level": {
                    "type": "integer",
                    "example": 2
                },
                "name": {
                    "type": "string",
                    "example": "test"
                }
            }
        },
        "auth.Login": {
            "type": "object",
            "required": [
                "account",
                "password"
            ],
            "properties": {
                "account": {
                    "type": "string",
                    "example": "admin"
                },
                "password": {
                    "type": "string",
                    "example": "qaz123"
                }
            }
        },
        "ipwhitelistings.IPWhitelistingModel": {
            "type": "object",
            "required": [
                "ip"
            ],
            "properties": {
                "enable": {
                    "type": "integer"
                },
                "ip": {
                    "type": "string",
                    "example": "127.0.0.1"
                },
                "remark": {
                    "type": "string"
                }
            }
        },
        "response.response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "message": {
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
	Version:     "1.0",
	Host:        "localhost:1002",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Golang Gin-Webcore API",
	Description: "This is a Gin-webcore",
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
