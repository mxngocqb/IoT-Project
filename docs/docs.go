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
        "/drivers": {
            "get": {
                "description": "Get all drivers from the database and returns them as JSON.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Driver"
                ],
                "summary": "Get all drivers",
                "responses": {
                    "200": {
                        "description": "List of drivers",
                        "schema": {
                            "$ref": "#/definitions/controller.DriverServerResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.DriverServerResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new driver",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Driver"
                ],
                "summary": "Create a new driver",
                "parameters": [
                    {
                        "description": "driver object that needs to be added",
                        "name": "driver",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Driver"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/drivers/{driverID}": {
            "get": {
                "description": "Retrieve a driver from the database by its ID and return it as JSON.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Driver"
                ],
                "summary": "Retrieve a driver by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Driver ID",
                        "name": "driverID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Driver information",
                        "schema": {
                            "$ref": "#/definitions/controller.DriverServerResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.DriverServerResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controller.DriverServerResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.DriverServerResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update information of a driver based on the provided data in JSON format.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Driver"
                ],
                "summary": "Update a driver",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Driver ID",
                        "name": "driverID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Driver object to be updated",
                        "name": "driver",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Driver"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Updated driver information",
                        "schema": {
                            "$ref": "#/definitions/controller.DriverServerResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.DriverServerResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.DriverServerResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a driver from the database based on its ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Driver"
                ],
                "summary": "Delete a driver",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Driver ID",
                        "name": "driverID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Driver deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/controller.DriverServerResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.DriverServerResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.DriverServerResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "status": {
                    "type": "string"
                }
            }
        },
        "model.Driver": {
            "type": "object",
            "required": [
                "driver_id",
                "driver_license"
            ],
            "properties": {
                "driver_id": {
                    "type": "integer"
                },
                "driver_license": {
                    "type": "string"
                },
                "driver_name": {
                    "type": "string"
                },
                "home_town": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "172.18.53.136:8080",
	BasePath:         "/",
	Schemes:          []string{"http", "https"},
	Title:            "IoT Proecjt",
	Description:      "This is a sample server for the IoT Project API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
