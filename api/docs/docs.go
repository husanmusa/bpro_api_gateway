// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://udevs.io",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/book": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "Get book list",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "category_id",
                        "name": "category_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GetBook ResponseBody",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/book_pro_service.GetBookListResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/book/": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This API for creating a new book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "Create a new book",
                "parameters": [
                    {
                        "description": "BookCreateRequest",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/book_pro_service.Book"
                        }
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/book/{book_id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "Get book by book_id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "book_id",
                        "name": "book_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GetBook ResponseBody",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/book_pro_service.Book"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "Update book by_id",
                "operationId": "update_book",
                "parameters": [
                    {
                        "description": "BookUpdateRequest",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/book_pro_service.Book"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success Update",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "Delete book by_id",
                "operationId": "delete_book",
                "parameters": [
                    {
                        "type": "string",
                        "description": "book_id",
                        "name": "book_id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success DeleteBook",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/book_Category/": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This API for creating a new bookCategory",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BookCategory"
                ],
                "summary": "Create a new bookCategory",
                "parameters": [
                    {
                        "description": "BookCategoryCreateRequest",
                        "name": "bookCategory",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/book_pro_service.BookCategory"
                        }
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/book_category": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BookCategory"
                ],
                "summary": "Get bookCategory list",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GetBookCategory ResponseBody",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/book_pro_service.GetBookCategoryListRes"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/book_category/{book_category_id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BookCategory"
                ],
                "summary": "Get bookCategory by bookCategory_id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bookCategory_id",
                        "name": "bookCategory_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GetBookCategory ResponseBody",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/book_pro_service.BookCategory"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BookCategory"
                ],
                "summary": "Update bookCategory by_id",
                "operationId": "update_bookCategory",
                "parameters": [
                    {
                        "description": "BookCategoryUpdateRequest",
                        "name": "bookCategory",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/book_pro_service.BookCategory"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success Update",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BookCategory"
                ],
                "summary": "Delete bookCategory by_id",
                "operationId": "delete_bookCategory",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bookCategory_id",
                        "name": "bookCategory_id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success DeleteBookCategory",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "book_pro_service.Book": {
            "type": "object",
            "properties": {
                "author_name": {
                    "type": "string"
                },
                "book_category": {
                    "type": "string"
                },
                "book_category_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "pages": {
                    "type": "integer"
                }
            }
        },
        "book_pro_service.BookCategory": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "book_pro_service.GetBookCategoryListRes": {
            "type": "object",
            "properties": {
                "book_categories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/book_pro_service.BookCategory"
                    }
                },
                "count": {
                    "type": "integer"
                }
            }
        },
        "book_pro_service.GetBookListResponse": {
            "type": "object",
            "properties": {
                "books": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/book_pro_service.Book"
                    }
                },
                "count": {
                    "type": "integer"
                }
            }
        },
        "http.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "description": {
                    "type": "string"
                },
                "error": {},
                "status": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
	Description:      "This is a api gateway",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
