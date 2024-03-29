// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
            "name": "João Da Silva Gomes",
            "url": "https://joaosilvagomes.pt",
            "email": "joao.s.gomes@outlook.pt"
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
        "/event": {
            "get": {
                "description": "Get all events from the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Get all events",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Event"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new event and add it to the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Create a new event",
                "parameters": [
                    {
                        "description": "Event object",
                        "name": "event",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Event"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Event"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponseModel"
                        }
                    }
                }
            }
        },
        "/event/{id}": {
            "get": {
                "description": "Get an event from the database by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Get an event by ID",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "Event ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Event"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponseModel"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an existing event in the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Update an existing event",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "Event ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Event object",
                        "name": "event",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Event"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Event"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponseModel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponseModel"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an event from the database by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Delete an event by ID",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "Event ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponseModel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponseModel"
                        }
                    }
                }
            }
        },
        "/events": {
            "get": {
                "description": "Retrieve all events from the repository and render the template with the events data",
                "consumes": [
                    "text/html"
                ],
                "produces": [
                    "text/html"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Render template for displaying all events",
                "responses": {
                    "200": {
                        "description": "HTML content with events data",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/events/{id}": {
            "get": {
                "description": "Retrieve the event specified by its ID from the repository and render the template with the event data",
                "consumes": [
                    "text/html"
                ],
                "produces": [
                    "text/html"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Render template for displaying a specific event",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "Event ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "HTML content with event data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/image": {
            "get": {
                "description": "Get all images from the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "images"
                ],
                "summary": "Get all images",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Image"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new image and add it to the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "images"
                ],
                "summary": "Create a new image",
                "parameters": [
                    {
                        "description": "Image object",
                        "name": "image",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Image"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Image"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponseModel"
                        }
                    }
                }
            }
        },
        "/image/{id}": {
            "get": {
                "description": "Get an image from the database by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "images"
                ],
                "summary": "Get an image by ID",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "Image ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Image"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponseModel"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an image from the database by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "images"
                ],
                "summary": "Delete an image by ID",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "Image ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponseModel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponseModel"
                        }
                    }
                }
            }
        },
        "/ws/event/{id}": {
            "get": {
                "description": "Establishes a WebSocket connection to get the remaining time until an event occurs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Get remaining time of an event via WebSocket",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "Event ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "models.ErrorResponseModel": {
            "type": "object",
            "properties": {
                "data": {},
                "date": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                }
            }
        },
        "models.Event": {
            "description": "Event information",
            "type": "object",
            "properties": {
                "capacity": {
                    "description": "The capacity of the event.\n@json:capacity",
                    "type": "integer"
                },
                "createdate": {
                    "description": "The creation date of the event.\n@json:createdate",
                    "type": "string"
                },
                "date": {
                    "description": "The date and time of the event.\n@json:date",
                    "type": "string"
                },
                "description": {
                    "description": "The description of the event.\n@json:description",
                    "type": "string"
                },
                "id": {
                    "description": "The unique identifier of the event.\n@json:id",
                    "type": "string"
                },
                "isactive": {
                    "description": "Indicates whether the event is active.\n@json:isactive",
                    "type": "boolean"
                },
                "location": {
                    "description": "The location of the event.\n@json:location",
                    "type": "string"
                },
                "name": {
                    "description": "The name of the event.\n@json:name",
                    "type": "string"
                },
                "src_image": {
                    "description": "The URL of the event image.\n@json:src_image",
                    "type": "string"
                },
                "temperature": {
                    "description": "The temperature of the event.\n@json:temperature",
                    "type": "number"
                },
                "versiondate": {
                    "description": "The version date of the event.\n@json:versiondate",
                    "type": "string"
                }
            }
        },
        "models.Image": {
            "description": "Image information",
            "type": "object",
            "properties": {
                "alt": {
                    "description": "Alt text for the image\n@json:alt",
                    "type": "string"
                },
                "description": {
                    "description": "Description of the image\n@json:description",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the image\n@json:id",
                    "type": "string"
                },
                "size": {
                    "description": "Size of the image in bytes\n@json:size",
                    "type": "integer"
                },
                "src_image": {
                    "description": "Source URL of the image\n@json:src_image",
                    "type": "string"
                }
            }
        }
    },
    "tags": [
        {
            "description": "MES4_PA_API",
            "name": "MES4_PA_API",
            "externalDocs": {
                "description": "https://github.com/joaosgomes/MES4_PA_API",
                "url": "https://github.com/joaosgomes/MES4_PA_API"
            }
        }
    ],
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:5000",
	BasePath:         "/",
	Schemes:          []string{"http", "https"},
	Title:            "MES4_PA_API",
	Description:      "Description MES4_PA_API\nDescription MES4_PA_API GO\nDescription MES4_PA_API GO Fiber",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
