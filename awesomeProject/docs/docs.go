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
        "/index": {
            "get": {
                "tags": [
                    "首页"
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
        },
        "/user/create_user": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "创建用户",
                "parameters": [
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "UID",
                        "name": "UID",
                        "in": "query",
                        "required": true
                    },
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "用户名",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "query",
                        "required": true
                    },
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "确认密码",
                        "name": "repassword",
                        "in": "query",
                        "required": true
                    },
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "query",
                        "required": true
                    },
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "phone",
                        "name": "phone",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/delete_user": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "UID",
                        "name": "UID",
                        "in": "query",
                        "required": true
                    },
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "用户名",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/update_user": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "修改用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "UID",
                        "name": "UID",
                        "in": "query",
                        "required": true
                    },
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "query",
                        "required": true
                    },
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "new name",
                        "name": "new_name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "new password",
                        "name": "new_password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
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
