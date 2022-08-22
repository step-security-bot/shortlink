// Package api GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package api

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Shortlink repository",
            "url": "https://github.com/batazor/shortlink/issues",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "MIT",
            "url": "http://www.opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/links": {
            "get": {
                "description": "List links",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "List links",
                "operationId": "list-links",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.ListResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update link",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update link",
                "operationId": "update-link",
                "parameters": [
                    {
                        "description": "Link",
                        "name": "link",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.UpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.UpdateResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Add link",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Add link",
                "operationId": "add-link",
                "parameters": [
                    {
                        "description": "Link",
                        "name": "link",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.AddRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.AddResponse"
                        }
                    }
                }
            }
        },
        "/links/{hash}": {
            "get": {
                "description": "Get link",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get link",
                "operationId": "get-link",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Hash",
                        "name": "hash",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.GetResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete link",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete link",
                "operationId": "delete-link",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Hash",
                        "name": "hash",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "github.com_batazor_shortlink_internal_services_link_domain_link_v1.Link": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "Create at",
                    "$ref": "#/definitions/timestamppb.Timestamp"
                },
                "describe": {
                    "description": "Describe of link",
                    "type": "string"
                },
                "hash": {
                    "description": "Hash by URL + salt",
                    "type": "string"
                },
                "updated_at": {
                    "description": "Update at",
                    "$ref": "#/definitions/timestamppb.Timestamp"
                },
                "url": {
                    "description": "URL",
                    "type": "string"
                }
            }
        },
        "github.com_batazor_shortlink_internal_services_link_infrastructure_rpc_link_v1.Link": {
            "type": "object",
            "properties": {
                "LinkServiceServer": {}
            }
        },
        "timestamppb.Timestamp": {
            "type": "object",
            "properties": {
                "nanos": {
                    "description": "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.",
                    "type": "integer"
                },
                "seconds": {
                    "description": "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.",
                    "type": "integer"
                }
            }
        },
        "v1.AddRequest": {
            "type": "object",
            "properties": {
                "link": {
                    "$ref": "#/definitions/github.com_batazor_shortlink_internal_services_link_domain_link_v1.Link"
                }
            }
        },
        "v1.AddResponse": {
            "type": "object",
            "properties": {
                "link": {
                    "$ref": "#/definitions/github.com_batazor_shortlink_internal_services_link_domain_link_v1.Link"
                }
            }
        },
        "v1.GetResponse": {
            "type": "object",
            "properties": {
                "link": {
                    "$ref": "#/definitions/github.com_batazor_shortlink_internal_services_link_domain_link_v1.Link"
                }
            }
        },
        "v1.Links": {
            "type": "object",
            "properties": {
                "link": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github.com_batazor_shortlink_internal_services_link_infrastructure_rpc_link_v1.Link"
                    }
                }
            }
        },
        "v1.ListResponse": {
            "type": "object",
            "properties": {
                "links": {
                    "$ref": "#/definitions/v1.Links"
                }
            }
        },
        "v1.UpdateRequest": {
            "type": "object",
            "properties": {
                "link": {
                    "$ref": "#/definitions/github.com_batazor_shortlink_internal_services_link_domain_link_v1.Link"
                }
            }
        },
        "v1.UpdateResponse": {
            "type": "object",
            "properties": {
                "link": {
                    "$ref": "#/definitions/github.com_batazor_shortlink_internal_services_link_domain_link_v1.Link"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:7070",
	BasePath:         "/api",
	Schemes:          []string{"http", "https"},
	Title:            "Shortlink API",
	Description:      "Shortlink API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
