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
        "/add-product": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Adds a product to the system. Only admins are allowed to use this function.",
                "consumes": [
                    "multipart/mixed"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "Add a product",
                "parameters": [
                    {
                        "description": "Product data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genprotos.ProductCReqForSwagger"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Product is added",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/delete-product/{id}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Deletes a product from the system. Only admins are allowed to use this function.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "Delete a product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id or email of the product",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "enum": [
                            "id",
                            "email"
                        ],
                        "type": "string",
                        "description": "Search with",
                        "name": "data",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Product is deleted",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/get-product/{id}": {
            "get": {
                "description": "Gets a product from the system.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "Get a product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id or email of the product",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "enum": [
                            "id",
                            "email"
                        ],
                        "type": "string",
                        "description": "Search with",
                        "name": "data",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Product data",
                        "schema": {
                            "$ref": "#/definitions/genprotos.ProductGRes"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/get-products": {
            "get": {
                "description": "Gets all products from the system.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "Get all products",
                "responses": {
                    "200": {
                        "description": "Product data",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/genprotos.ProductGARes"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/update-product-count/{product_id}": {
            "put": {
                "description": "Updates the count of a product.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "Update product count",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "product_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Count data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genprotos.ProductCountUReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Product count updated",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/update-product-image/{id}": {
            "put": {
                "description": "Updates the image URL of a product.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "Update product image URL",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Image URL data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genprotos.ProductImageUReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Product image URL updated",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/update-product-rating/{product_id}": {
            "put": {
                "description": "Updates the rating of a product.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "Update product rating",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "product_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Rating data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genprotos.ProductRatingUReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Product rating updated",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/update-product/{id}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Updates a product in the system. Only admins are allowed to use this function.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "Update a product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id or email of the product",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "enum": [
                            "id",
                            "email"
                        ],
                        "type": "string",
                        "description": "Search with",
                        "name": "data",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Product data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genprotos.ProductUReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Product is updated",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "genprotos.ProductCReqForSwagger": {
            "type": "object",
            "properties": {
                "additional_details": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "category": {
                    "type": "string"
                },
                "count": {
                    "type": "number"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "seller": {
                    "type": "string"
                },
                "weight": {
                    "type": "number"
                }
            }
        },
        "genprotos.ProductCountUReq": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "number"
                },
                "product_id": {
                    "type": "string"
                }
            }
        },
        "genprotos.ProductGARes": {
            "type": "object",
            "properties": {
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/genprotos.ProductGRes"
                    }
                }
            }
        },
        "genprotos.ProductGRes": {
            "type": "object",
            "properties": {
                "additional_details": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "category": {
                    "type": "string"
                },
                "count": {
                    "type": "number"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "img_url": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "rating": {
                    "type": "number"
                },
                "seller": {
                    "type": "string"
                },
                "weight": {
                    "type": "number"
                }
            }
        },
        "genprotos.ProductImageUReq": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "img_url": {
                    "type": "string"
                }
            }
        },
        "genprotos.ProductRatingUReq": {
            "type": "object",
            "properties": {
                "product_id": {
                    "type": "string"
                },
                "rate": {
                    "type": "number"
                }
            }
        },
        "genprotos.ProductUReq": {
            "type": "object",
            "properties": {
                "additional_details": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "category": {
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
                "seller": {
                    "type": "string"
                },
                "weight": {
                    "type": "number"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Swaggers of Product manager",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
