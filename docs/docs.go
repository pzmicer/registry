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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/addService": {
            "post": {
                "parameters": [
                    {
                        "description": "Body parameter",
                        "name": "service",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.ServiceNoID"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/checkService": {
            "get": {
                "parameters": [
                    {
                        "type": "string",
                        "description": "Service name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Example \\n {\"result\": true\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/getServiceInfo": {
            "get": {
                "parameters": [
                    {
                        "type": "string",
                        "description": "Service name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Service"
                        }
                    }
                }
            }
        },
        "/getServiceList": {
            "get": {
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Service"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Service": {
            "type": "object",
            "properties": {
                "cost": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "key_method": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "main.ServiceNoID": {
            "type": "object",
            "properties": {
                "cost": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer"
                },
                "key_method": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "url": {
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
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
	swag.Register("swagger", &s{})
}
