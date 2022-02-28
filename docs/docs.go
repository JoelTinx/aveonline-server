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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "fiber@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/invoices": {
            "get": {
                "description": "get the list of invoices.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "invoices"
                ],
                "summary": "Show the list of invoices.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Invoice"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "create a new invoice.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "invoices"
                ],
                "summary": "create a new invoice.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Invoice"
                        }
                    }
                }
            }
        },
        "/api/medicines": {
            "get": {
                "description": "get the list of medicines.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "medicines"
                ],
                "summary": "Show the list of medicines.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Medicine"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "create a new medicine.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "medicines"
                ],
                "summary": "create a new medicine.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Medicine"
                        }
                    }
                }
            }
        },
        "/api/promotions": {
            "get": {
                "description": "get the list of promotions.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "promotions"
                ],
                "summary": "Show the list of promotions.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Promotion"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "create a new promotion.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "promotions"
                ],
                "summary": "create a new promotion.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Promotion"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Invoice": {
            "type": "object",
            "properties": {
                "creation_date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "total": {
                    "type": "number"
                }
            }
        },
        "model.Medicine": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "model.Promotion": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "end_date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "percentage": {
                    "type": "number"
                },
                "start_date": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:5000",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "API Facturacion",
	Description:      "Permite gestionar la facturación de Aveonline.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
