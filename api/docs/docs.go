// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "GNU General Public License v3.0",
            "url": "https://github.com/dreamvo/gilfoyle/blob/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/healthz": {
            "get": {
                "description": "Check for the health of the service",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Instance"
                ],
                "summary": "Check service status",
                "operationId": "checkHealth",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.DataResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/api.HealthCheckResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/medias": {
            "get": {
                "description": "Get latest created medias",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Medias"
                ],
                "summary": "Query medias",
                "operationId": "getAllMedias",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Max number of results",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Number of results to ignore",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.DataResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/ent.Media"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new media",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Medias"
                ],
                "summary": "Create a media",
                "operationId": "createMedia",
                "parameters": [
                    {
                        "description": "Media data",
                        "name": "media",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.CreateMedia"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.DataResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/ent.Media"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/medias/{id}": {
            "get": {
                "description": "Get one media",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Medias"
                ],
                "summary": "Get a media",
                "operationId": "getMedia",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Media ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.DataResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/ent.Media"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete one media",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Medias"
                ],
                "summary": "Delete a media",
                "operationId": "deleteMedia",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Media ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.DataResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update an existing media",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Medias"
                ],
                "summary": "Update a media",
                "operationId": "updateMedia",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Media ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Media data",
                        "name": "media",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.UpdateMedia"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.DataResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/ent.Media"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/medias/{id}/upload/audio": {
            "post": {
                "description": "Upload a new audio file for a given media ID",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Medias"
                ],
                "summary": "Upload a audio file",
                "operationId": "uploadAudio",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Media identifier",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Audio file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.DataResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/api.FileFormat"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/medias/{id}/upload/video": {
            "post": {
                "description": "Upload a new video file for a given media ID",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Medias"
                ],
                "summary": "Upload a video file",
                "operationId": "uploadVideo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Media identifier",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Video file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.DataResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/api.FileFormat"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/medias/{media_id}/attachments": {
            "get": {
                "description": "Get attachments of a media",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Attachments"
                ],
                "summary": "Get attachments of a media",
                "operationId": "getMediaAttachments",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Media identifier",
                        "name": "media_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.DataResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Add attachment to a media",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Attachments"
                ],
                "summary": "Add attachment to a media",
                "operationId": "addMediaAttachment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Media identifier",
                        "name": "media_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Attachment file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "description": "Attachment metadata",
                        "name": "attachment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.AddMediaAttachment"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.DataResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/medias/{media_id}/attachments/{attachment_id}": {
            "delete": {
                "description": "Delete attachment of a media",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Attachments"
                ],
                "summary": "Delete attachment of a media",
                "operationId": "deleteMediaAttachment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Media identifier",
                        "name": "media_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Attachment identifier",
                        "name": "attachment_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.DataResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/medias/{media_id}/stream/{preset}": {
            "get": {
                "description": "Get stream from media file",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stream"
                ],
                "summary": "Get stream from media file",
                "operationId": "streamMedia",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Media identifier",
                        "name": "media_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Encoder preset",
                        "name": "preset",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.DataResponse"
                        },
                        "headers": {
                            "Content-Type": {
                                "type": "string",
                                "description": "application/octet-stream"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/metricsz": {
            "get": {
                "description": "Get metrics about this Gilfoyle instance",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Instance"
                ],
                "summary": "Get instance metrics",
                "operationId": "getMetrics",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.DataResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.AddMediaAttachment": {
            "type": "object",
            "required": [
                "key"
            ],
            "properties": {
                "key": {
                    "type": "string",
                    "example": "subtitle_fr_FR"
                }
            }
        },
        "api.CreateMedia": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "title": {
                    "type": "string",
                    "example": "Sheep Discovers How To Use A Trampoline"
                }
            }
        },
        "api.FileFormat": {
            "type": "object",
            "properties": {
                "bit_rate": {
                    "type": "string"
                },
                "duration": {
                    "type": "string",
                    "example": "0"
                },
                "filename": {
                    "type": "string"
                },
                "format_long_name": {
                    "type": "string"
                },
                "format_name": {
                    "type": "string"
                },
                "nb_programs": {
                    "type": "integer"
                },
                "nb_streams": {
                    "type": "integer"
                },
                "probe_score": {
                    "type": "integer"
                },
                "size": {
                    "type": "string"
                },
                "start_time": {
                    "type": "string",
                    "example": "0"
                }
            }
        },
        "api.HealthCheckResponse": {
            "type": "object",
            "properties": {
                "commit": {
                    "type": "string"
                },
                "tag": {
                    "type": "string"
                }
            }
        },
        "api.UpdateMedia": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "title": {
                    "type": "string",
                    "example": "Sheep Discovers How To Use A Trampoline"
                }
            }
        },
        "ent.Media": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "edges": {
                    "type": "MediaEdges"
                },
                "id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "util.DataResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "data": {
                    "type": "object"
                }
            }
        },
        "util.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "v1",
	Host:        "demo-v1.gilfoyle.dreamvo.com",
	BasePath:    "/",
	Schemes:     []string{"http", "https"},
	Title:       "Gilfoyle server",
	Description: "Cloud-native media hosting & streaming server for businesses.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
