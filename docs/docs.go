// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-10-13 10:51:23.5516536 +0800 CST m=+0.129920301

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is a demo server.",
        "title": "RESTful API DEMO",
        "contact": {
            "name": "Muninn",
            "email": "hyacinthus@gmail.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://github.com/hyacinthus/restdemo/blob/master/LICENSE"
        },
        "version": "1.0"
    },
    "host": "demo.crandom.com",
    "basePath": "/",
    "paths": {
        "/login": {
            "post": {
                "description": "用户登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "description": "登录凭证",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.Token"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.httpError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.httpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.httpError"
                        }
                    }
                }
            }
        },
        "/notes": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取用户的全部笔记，有分页，默认一页10条。",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "笔记"
                ],
                "summary": "获取用户笔记列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页几条",
                        "name": "per_page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Note"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.httpError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.httpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.httpError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "新建一条笔记",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "笔记"
                ],
                "summary": "新建笔记",
                "parameters": [
                    {
                        "description": "笔记内容",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.Note"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.Note"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.httpError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.httpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.httpError"
                        }
                    }
                }
            }
        },
        "/notes/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取指定id的笔记",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "笔记"
                ],
                "summary": "获取笔记",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "笔记编号",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.Note"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.httpError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.httpError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.httpError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.httpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.httpError"
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
                "description": "更新指定id的笔记",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "笔记"
                ],
                "summary": "更新笔记",
                "parameters": [
                    {
                        "description": "更新内容",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.NoteUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.Note"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.httpError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.httpError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.httpError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.httpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.httpError"
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
                "description": "删除指定id的笔记",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "笔记"
                ],
                "summary": "删除笔记",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "笔记编号",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.httpError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.httpError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.httpError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.httpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.httpError"
                        }
                    }
                }
            }
        },
        "/public/notes": {
            "get": {
                "description": "获取公开的全部笔记，有分页，默认一页10条。",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "笔记"
                ],
                "summary": "获取公开笔记列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页几条",
                        "name": "per_page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Note"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.httpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.httpError"
                        }
                    }
                }
            }
        },
        "/public/notes/{id}": {
            "get": {
                "description": "获取指定id的公开笔记",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "笔记"
                ],
                "summary": "获取公开笔记",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "笔记编号",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.Note"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.httpError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.httpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.httpError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.FileURL": {
            "type": "object"
        },
        "main.LoginRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "description": "密码\n",
                    "type": "string"
                },
                "username": {
                    "description": "用户名\n",
                    "type": "string"
                }
            }
        },
        "main.Note": {
            "type": "object",
            "properties": {
                "content": {
                    "description": "内容\n",
                    "type": "string"
                },
                "created_at": {
                    "description": "创建时间\n",
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "description": "题图\n",
                    "type": "object",
                    "$ref": "#/definitions/main.FileURL"
                },
                "is_public": {
                    "description": "是否公开\n",
                    "type": "boolean"
                },
                "title": {
                    "description": "标题\n",
                    "type": "string"
                },
                "updated_at": {
                    "description": "最后更新时间\n",
                    "type": "string"
                },
                "user_id": {
                    "description": "所属用户\n",
                    "type": "integer"
                }
            }
        },
        "main.NoteUpdate": {
            "type": "object",
            "properties": {
                "content": {
                    "description": "内容\n",
                    "type": "string"
                },
                "image": {
                    "description": "题图\n",
                    "type": "object",
                    "$ref": "#/definitions/main.FileURL"
                },
                "is_public": {
                    "description": "是否公开\n",
                    "type": "boolean"
                },
                "title": {
                    "description": "标题\n",
                    "type": "string"
                }
            }
        },
        "main.Token": {
            "type": "object",
            "properties": {
                "expires_at": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "main.httpError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "error": {
                    "description": "错误代码，为英文字符串，前端可用此判断大的错误类型。\n",
                    "type": "string"
                },
                "message": {
                    "description": "错误消息，为详细错误描述，前端可选择性的展示此字段。\n",
                    "type": "string"
                }
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}
