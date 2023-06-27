// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/api/v1/bahan": {
            "post": {
                "description": "Create a new master data bahan.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bahan"
                ],
                "summary": "Create bahan",
                "parameters": [
                    {
                        "description": "Create bahan",
                        "name": "bahan",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.BahanInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/bahan/delete/{id}": {
            "delete": {
                "description": "Return data boolean.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bahan"
                ],
                "summary": "Delete bahan by id.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "delete bahan by id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/bahan/{id}": {
            "get": {
                "description": "Return data master bahan where similar with id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bahan"
                ],
                "summary": "Get Single bahan by id.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "get bahan by id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            },
            "put": {
                "description": "Update data master bahan.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bahan"
                ],
                "summary": "Update bahan",
                "parameters": [
                    {
                        "description": "Update bahan",
                        "name": "bahan",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.BahanInput"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "find bahan by id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/kategori": {
            "post": {
                "description": "Create a new master data kategori.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "kategori"
                ],
                "summary": "Create kategori",
                "parameters": [
                    {
                        "description": "Create kategori",
                        "name": "kategori",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.KategoriInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/kategori/delete/{id}": {
            "delete": {
                "description": "Return data boolean.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "kategori"
                ],
                "summary": "Delete kategori by id.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "delete kategori by id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/kategori/{id}": {
            "get": {
                "description": "Return data master kategori where similar with id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "kategori"
                ],
                "summary": "Get Single kategori by id.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "get kategori by id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            },
            "put": {
                "description": "Update data master kategori.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "kategori"
                ],
                "summary": "Update kategori",
                "parameters": [
                    {
                        "description": "Update kategori",
                        "name": "bahan",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.KategoriInput"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "find kategori by id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.BahanInput": {
            "type": "object",
            "required": [
                "nama"
            ],
            "properties": {
                "nama": {
                    "type": "string"
                }
            }
        },
        "models.KategoriInput": {
            "type": "object",
            "required": [
                "nama"
            ],
            "properties": {
                "nama": {
                    "type": "string"
                }
            }
        },
        "utils.Meta": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "utils.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "meta": {
                    "$ref": "#/definitions/utils.Meta"
                }
            }
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}