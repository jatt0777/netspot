// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "asr@nabla.ovh"
        },
        "license": {
            "name": "GPLv3",
            "url": "https://www.gnu.org/licenses/gpl-3.0.en.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/config": {
            "get": {
                "description": "You can fetch the netspot config through this endpoint",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get the config of the IDS",
                "responses": {
                    "200": {
                        "description": "Acknowledge message",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/api.apiError"
                        }
                    },
                    "405": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/api.apiError"
                        }
                    },
                    "500": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/api.apiError"
                        }
                    }
                }
            },
            "post": {
                "description": "You can update the netspot config through this endpoint",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain",
                    "application/json"
                ],
                "summary": "Update the config of the IDS",
                "responses": {
                    "201": {
                        "description": "Acknowledge message",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/api.apiError"
                        }
                    },
                    "405": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/api.apiError"
                        }
                    },
                    "500": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/api.apiError"
                        }
                    }
                }
            }
        },
        "/api/devices": {
            "get": {
                "description": "This returns the list of the network interfaces that can be monitored",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "List the available devices",
                "responses": {
                    "200": {
                        "description": "list of the available devices",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "error message",
                        "schema": {
                            "$ref": "#/definitions/api.apiError"
                        }
                    }
                }
            }
        },
        "/api/ping": {
            "get": {
                "description": "This endpoints basically aims to check if the server is up",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Server healthcheck",
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "405": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/api.apiError"
                        }
                    }
                }
            }
        },
        "/api/run": {
            "post": {
                "description": "Use this path to start/stop the IDS",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Manage the IDS status",
                "parameters": [
                    {
                        "enum": [
                            "\"start\"",
                            "\"stop\""
                        ],
                        "description": "the action to perform",
                        "name": "action",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Comment about the action performed",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/api.apiError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.apiError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Oh my god! Something wrong happened"
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
	Version:     "2.0",
	Host:        "localhost:11000",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Netspot API",
	Description: "Netspot as a service",
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
