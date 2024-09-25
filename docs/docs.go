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
        "/auth/profile": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "api to get profile user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "GetProfile",
                "parameters": [
                    {
                        "enum": [
                            "WEB",
                            "IPHONE",
                            "ANDROID"
                        ],
                        "type": "string",
                        "name": "device",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SwaggoGetProfileResSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.SwaggoGetProfileResBadRequest"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.SwaggoGetProfileResUnauthorized"
                        }
                    }
                }
            }
        },
        "/auth/signin": {
            "post": {
                "description": "sign in api",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Signin",
                "parameters": [
                    {
                        "description": "Request Body signin",
                        "name": "PostSignin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqPostSignin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SwaggoSigninResSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.SwaggoSigninResBadRequest"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.SwaggoSigninResNotFound"
                        }
                    }
                }
            }
        },
        "/notes": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "api to get all notes",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notes"
                ],
                "summary": "Get All Notes",
                "parameters": [
                    {
                        "type": "integer",
                        "example": 0,
                        "name": "all",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "title",
                        "name": "keyword",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 10,
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 0,
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "title",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "asc",
                            "desc"
                        ],
                        "type": "string",
                        "example": "asc",
                        "name": "order_by",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SwaggoGetNotesResSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.SwaggoGetNotesResBadRequest"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.SwaggoGetNotesResUnauthorized"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.SwaggoGetNotesResNotFound"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "api to update notes",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notes"
                ],
                "summary": "Put Notes",
                "parameters": [
                    {
                        "description": "Request Param Put Notes",
                        "name": "PutNotes",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqPutNotes"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SwaggoPutNotesResSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.SwaggoPutNotesResBadRequest"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.SwaggoPutNotesResUnauthorized"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "api to create notes",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notes"
                ],
                "summary": "Post Notes",
                "parameters": [
                    {
                        "description": "Request Param Post Notes",
                        "name": "PostNotes",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqPostNotes"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SwaggoPostNotesResSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.SwaggoPostNotesResBadRequest"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.SwaggoPostNotesResUnauthorized"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "api to delete notes",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notes"
                ],
                "summary": "Delete Notes",
                "parameters": [
                    {
                        "description": "Request Param Delete Notes",
                        "name": "DeleteNotes",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqDeleteNotes"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SwaggoDeleteNotesResSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.SwaggoDeleteNotesResBadRequest"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.SwaggoDeleteNotesResUnauthorized"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.EmptyResponse": {
            "type": "object"
        },
        "models.NotesData": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "timestamp"
                },
                "description": {
                    "type": "string",
                    "example": "description"
                },
                "title": {
                    "type": "string",
                    "example": "title"
                },
                "uid": {
                    "type": "string",
                    "example": "uuid"
                }
            }
        },
        "models.ReqDeleteNotes": {
            "type": "object",
            "required": [
                "uid_notes"
            ],
            "properties": {
                "uid_notes": {
                    "type": "string"
                }
            }
        },
        "models.ReqPostNotes": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.ReqPostSignin": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                }
            }
        },
        "models.ReqPutNotes": {
            "type": "object",
            "required": [
                "title",
                "uid_notes"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "uid_notes": {
                    "type": "string"
                }
            }
        },
        "models.ResSignin": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string",
                    "example": "token"
                }
            }
        },
        "models.SwaggoDeleteNotesResBadRequest": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.EmptyResponse"
                },
                "message": {
                    "type": "string",
                    "example": "valid param required"
                },
                "status": {
                    "type": "boolean",
                    "example": false
                },
                "status_code": {
                    "type": "integer",
                    "example": 400
                },
                "total_data": {
                    "type": "integer",
                    "example": 0
                }
            }
        },
        "models.SwaggoDeleteNotesResSuccess": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.EmptyResponse"
                },
                "message": {
                    "type": "string",
                    "example": "success"
                },
                "status": {
                    "type": "boolean",
                    "example": true
                },
                "status_code": {
                    "type": "integer",
                    "example": 200
                },
                "total_data": {
                    "type": "integer",
                    "example": 0
                }
            }
        },
        "models.SwaggoDeleteNotesResUnauthorized": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.EmptyResponse"
                },
                "message": {
                    "type": "string",
                    "example": "token is required"
                },
                "status": {
                    "type": "boolean",
                    "example": false
                },
                "status_code": {
                    "type": "integer",
                    "example": 401
                },
                "total_data": {
                    "type": "integer",
                    "example": 0
                }
            }
        },
        "models.SwaggoGetNotesResBadRequest": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.EmptyResponse"
                },
                "message": {
                    "type": "string",
                    "example": "valid param required"
                },
                "status": {
                    "type": "boolean",
                    "example": false
                },
                "status_code": {
                    "type": "integer",
                    "example": 400
                },
                "total_data": {
                    "type": "integer",
                    "example": 0
                }
            }
        },
        "models.SwaggoGetNotesResNotFound": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.EmptyResponse"
                },
                "message": {
                    "type": "string",
                    "example": "data not found"
                },
                "status": {
                    "type": "boolean",
                    "example": false
                },
                "status_code": {
                    "type": "integer",
                    "example": 400
                },
                "total_data": {
                    "type": "integer",
                    "example": 0
                }
            }
        },
        "models.SwaggoGetNotesResSuccess": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.NotesData"
                    }
                },
                "message": {
                    "type": "string",
                    "example": "success"
                },
                "status": {
                    "type": "boolean",
                    "example": true
                },
                "status_code": {
                    "type": "integer",
                    "example": 200
                },
                "total_data": {
                    "type": "integer",
                    "example": 0
                }
            }
        },
        "models.SwaggoGetNotesResUnauthorized": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.EmptyResponse"
                },
                "message": {
                    "type": "string",
                    "example": "token is required"
                },
                "status": {
                    "type": "boolean",
                    "example": false
                },
                "status_code": {
                    "type": "integer",
                    "example": 401
                },
                "total_data": {
                    "type": "integer",
                    "example": 0
                }
            }
        },
        "models.SwaggoGetProfileResBadRequest": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.EmptyResponse"
                },
                "message": {
                    "type": "string",
                    "example": "valid param required"
                },
                "status": {
                    "type": "boolean",
                    "example": false
                },
                "status_code": {
                    "type": "integer",
                    "example": 400
                },
                "total_data": {
                    "type": "integer",
                    "example": 0
                }
            }
        },
        "models.SwaggoGetProfileResSuccess": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/starterapi_modules_auth_models.UserData"
                },
                "message": {
                    "type": "string",
                    "example": "success"
                },
                "status": {
                    "type": "boolean",
                    "example": true
                },
                "status_code": {
                    "type": "integer",
                    "example": 200
                },
                "total_data": {
                    "type": "integer",
                    "example": 0
                }
            }
        },
        "models.SwaggoGetProfileResUnauthorized": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.EmptyResponse"
                },
                "message": {
                    "type": "string",
                    "example": "data not found, token is required"
                },
                "status": {
                    "type": "boolean",
                    "example": false
                },
                "status_code": {
                    "type": "integer",
                    "example": 401
                },
                "total_data": {
                    "type": "integer",
                    "example": 0
                }
            }
        },
        "models.SwaggoPostNotesResBadRequest": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.EmptyResponse"
                },
                "message": {
                    "type": "string",
                    "example": "valid param required"
                },
                "status": {
                    "type": "boolean",
                    "example": false
                },
                "status_code": {
                    "type": "integer",
                    "example": 400
                },
                "total_data": {
                    "type": "integer",
                    "example": 0
                }
            }
        },
        "models.SwaggoPostNotesResSuccess": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.EmptyResponse"
                },
                "message": {
                    "type": "string",
                    "example": "success"
                },
                "status": {
                    "type": "boolean",
                    "example": true
                },
                "status_code": {
                    "type": "integer",
                    "example": 200
                },
                "total_data": {
                    "type": "integer",
                    "example": 0
                }
            }
        },
        "models.SwaggoPostNotesResUnauthorized": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.EmptyResponse"
                },
                "message": {
                    "type": "string",
                    "example": "token is required"
                },
                "status": {
                    "type": "boolean",
                    "example": false
                },
                "status_code": {
                    "type": "integer",
                    "example": 401
                },
                "total_data": {
                    "type": "integer",
                    "example": 0
                }
            }
        },
        "models.SwaggoPutNotesResBadRequest": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.EmptyResponse"
                },
                "message": {
                    "type": "string",
                    "example": "valid param required"
                },
                "status": {
                    "type": "boolean",
                    "example": false
                },
                "status_code": {
                    "type": "integer",
                    "example": 400
                },
                "total_data": {
                    "type": "integer",
                    "example": 0
                }
            }
        },
        "models.SwaggoPutNotesResSuccess": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.EmptyResponse"
                },
                "message": {
                    "type": "string",
                    "example": "success"
                },
                "status": {
                    "type": "boolean",
                    "example": true
                },
                "status_code": {
                    "type": "integer",
                    "example": 200
                },
                "total_data": {
                    "type": "integer",
                    "example": 0
                }
            }
        },
        "models.SwaggoPutNotesResUnauthorized": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.EmptyResponse"
                },
                "message": {
                    "type": "string",
                    "example": "token is required"
                },
                "status": {
                    "type": "boolean",
                    "example": false
                },
                "status_code": {
                    "type": "integer",
                    "example": 401
                },
                "total_data": {
                    "type": "integer",
                    "example": 0
                }
            }
        },
        "models.SwaggoSigninResBadRequest": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.EmptyResponse"
                },
                "message": {
                    "type": "string",
                    "example": "valid param required, password missmatch"
                },
                "status": {
                    "type": "boolean",
                    "example": false
                },
                "status_code": {
                    "type": "integer",
                    "example": 400
                },
                "total_data": {
                    "type": "integer",
                    "example": 0
                }
            }
        },
        "models.SwaggoSigninResNotFound": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.EmptyResponse"
                },
                "message": {
                    "type": "string",
                    "example": "error get user"
                },
                "status": {
                    "type": "boolean",
                    "example": false
                },
                "status_code": {
                    "type": "integer",
                    "example": 404
                },
                "total_data": {
                    "type": "integer",
                    "example": 0
                }
            }
        },
        "models.SwaggoSigninResSuccess": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.ResSignin"
                },
                "message": {
                    "type": "string",
                    "example": "success"
                },
                "status": {
                    "type": "boolean",
                    "example": true
                },
                "status_code": {
                    "type": "integer",
                    "example": 200
                },
                "total_data": {
                    "type": "integer",
                    "example": 0
                }
            }
        },
        "starterapi_modules_auth_models.UserData": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "example@mail.com"
                },
                "uid": {
                    "type": "string",
                    "example": "uid"
                },
                "username": {
                    "type": "string",
                    "example": "user_example"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "Type \"Bearer\" followed by a space and JWT token.",
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
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
