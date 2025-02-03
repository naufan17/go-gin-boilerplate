// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `
{
	"schemes": {{ marshal .Schemes }},
	"swagger": "2.0",
	"info": {
		"title": "{{.Title}}",
		"version": "{{.Version}}",
		"description": "{{escape .Description}}",
	},
	"host": "{{.Host}}",
	"basePath": "{{.BasePath}}",
	"tags": [
		{
			"name": "Auth",
			"description": "Authentication management"
		},
		{
			"name": "Account",
			"description": "Account management"
		}
	],
	"components": {
		"securitySchemes": {
			"BearerAuth": {
				"type": "http",
				"scheme": "bearer",
				"bearerFormat": "JWT"
			}
		}
	},
	"paths": {
		"/api/v1/auth/register": {
			"post": {
				"summary": "Create new account",
				"tags": ["Auth"],
				"parameters": [
					{
						"name": "name",
						"in": "body",
						"required": true,
						"schema": {
							"type": "string"
						}
					},
					{
						"name": "email",
						"in": "body",
						"required": true,
						"schema": {
							"type": "string"
						}
					},
					{
						"name": "password",
						"in": "body",
						"required": true,
						"schema": {
							"type": "string"
						}
					},
					{
						"name": "confirm_password",
						"in": "body",
						"required": true,
						"schema": {
							"type": "string"
						}
					}
				],
				"produces": [
					"application/json"
				],
				"responses": {
					"201": {
						"description": "Account created successfully",
						"schema": {
							"type": "object",
							"properties": {
								"message": {
									"type": "string"
								}
							}
						}
					},
					"400": {
						"description": "Bad request",
						"schema": {
							"type": "object",
							"properties": {
								"error": {
									"type": "string"
								}
							}
						}
					},
					"409": {
						"description": "Email already exists",
						"schema": {
							"type": "object",
							"properties": {
								"error": {
									"type": "string"
								}
							}
						}
					},
					"500": {
						"description": "Error creating account",
						"schema": {
							"type": "object",
							"properties": {
								"error": {
									"type": "string"
								}
							}
						}
					}
				}
			}
		},
		"/api/v1/auth/login": {
			"post": {
				"summary": "Login to existing account",
				"tags": ["Auth"],
				"parameters": [
					{
						"name": "email",
						"in": "body",
						"required": true,
						"schema": {
							"type": "string"
						}
					},
					{
						"name": "password",
						"in": "body",
						"required": true,
						"schema": {
							"type": "string"
						}
					}
				],
				"produces": [
					"application/json"
				],
				"responses": {
					"200": {
						"description": "Login successful",
						"schema": {
							"type": "object",
							"properties": {
								"data": {
									"type": "object",
									"properties": {
										"accessToken": {
											"type": "string"
										},
										"expiresIn": {
											"type": "number"
										},
										"tokenType": {
											"type": "string"
										}
									}
								}
							}
						}
					},
					"400": {
						"description": "Bad request",
						"schema": {
							"type": "object",
							"properties": {
								"error": {
									"type": "string"
								}
							}
						}
					},
					"401": {
						"description": "Invalid email or password",
						"schema": {
							"type": "object",
							"properties": {
								"error": {
									"type": "string"
								}
							}
						}			
					},
					"500": {
						"description": "Error logging in",
						"schema": {
							"type": "object",
							"properties": {
								"error": {
									"type": "string"
								}
							}
						}
					}
				}
			}
		},
		"/api/v1/account/profile": {
			"get": {
				"summary": "Get current account profile",
				"tags": ["Account"],
				"parameters": [
					{
						"name": "Authorization",
						"in": "header",
						"required": true,
						"schema": {
							"type": "string"
						}
					}
				],
				"produces": [
					"application/json"
				],
				"responses": {
					"200": {
						"description": "Get current users",
						"schema": {
							"type": "object",
							"properties": {
								"data": {
									"type": "object",
									"properties": {
										"id": {
											"type": "number"
										},
										"name": {
											"type": "string"
										},
										"email": {
											"type": "string"
										}
									}
								}
							}
						}
					},
					"401": {
						"description": "Invalid token",
						"schema": {
							"type": "object",
							"properties": {
								"error": {
									"type": "string"
								}
							}
						}
					},
					"500": {
						"description": "Error getting users",
						"schema": {
							"type": "object",
							"properties": {
								"error": {
									"type": "string"
								}
							}
						}
					}
				}
			}
		}
	}
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Go Rest API Boilerplate",
	Description:      "Minimalist project structure using Gin to build REST API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
