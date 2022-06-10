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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/check": {
            "get": {
                "description": "登录态校验",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "登录态校验"
                ],
                "summary": "登录态校验",
                "parameters": [
                    {
                        "type": "string",
                        "description": "b端用户token",
                        "name": "b_access_token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.AuthStatus"
                        }
                    }
                }
            }
        },
        "/signin/username": {
            "post": {
                "description": "用户名登录",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    ""
                ],
                "summary": "用户名登录",
                "parameters": [
                    {
                        "description": "注释",
                        "name": "xxx",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/request.SignInUsernameRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.SignInUsernameRsp"
                        }
                    }
                }
            }
        },
        "/signup/username": {
            "post": {
                "description": "用户名注册",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    ""
                ],
                "summary": "用户名注册",
                "parameters": [
                    {
                        "description": "注释",
                        "name": "xxx",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/request.SignUpUsernameRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.SignUpRsp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.AuthStatus": {
            "type": "object",
            "properties": {
                "app_id": {
                    "type": "string"
                },
                "expire_at": {
                    "description": "过期时间",
                    "type": "string",
                    "example": "2022-05-16 23:00:00"
                },
                "uid": {
                    "description": "b端用户id",
                    "type": "string",
                    "example": "u_12345678901"
                }
            }
        },
        "request.SignInUsernameRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "description": "密码，需要包含大小写数字和特殊字符",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string",
                    "example": "zhuyan"
                }
            }
        },
        "request.SignUpUsernameRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "description": "密码，需要包含大小写数字和特殊字符",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string",
                    "example": "zhuyan"
                }
            }
        },
        "response.SignInUsernameRsp": {
            "type": "object",
            "properties": {
                "b_access_token": {
                    "description": "b端用户token",
                    "type": "string",
                    "example": "b_u_uasdasd"
                },
                "uid": {
                    "description": "用户ID",
                    "type": "string"
                }
            }
        },
        "response.SignUpRsp": {
            "type": "object",
            "properties": {
                "b_access_token": {
                    "description": "b端用户token",
                    "type": "string",
                    "example": "b_u_uasdasd"
                },
                "uid": {
                    "description": "用户ID",
                    "type": "string"
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
