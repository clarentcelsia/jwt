// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Developer"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user/customer": {
            "post": {
                "description": "do ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "example"
                ],
                "summary": "ping example",
                "parameters": [
                    {
                        "description": "customer",
                        "name": "Data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Customer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Base": {
            "type": "object",
            "properties": {
                "crt_at": {
                    "type": "string"
                },
                "upd_at": {
                    "type": "string"
                }
            }
        },
        "model.Customer": {
            "type": "object",
            "properties": {
                "base": {
                    "$ref": "#/definitions/model.Base"
                },
                "customer_address": {
                    "type": "string"
                },
                "customer_dob": {
                    "type": "string"
                },
                "customer_email": {
                    "type": "string"
                },
                "customer_id": {
                    "type": "string"
                },
                "customer_name": {
                    "type": "string"
                },
                "customer_phone": {
                    "type": "string"
                },
                "is_deleted": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/api/v1/restaurant",
	Schemes:          []string{},
	Title:            "Web Order API",
	Description:      "This page is API documentation for all services relating common data or operation.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
