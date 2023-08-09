// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "http://www.parodos.dev",
            "email": "parodos@redhat.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/groups": {
            "get": {
                "description": "return the list of groups registered.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get a list of groups",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/pkg_workflows.Group"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/groups/{group_id}": {
            "get": {
                "description": "return the details of a given registered group.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get the details of a registered group",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/pkg_workflows.GroupDetails"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/groups/{group_id}/workflows": {
            "get": {
                "description": "return the list of workflows definitions registered in the given group.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get a list of workflows definitions in the group",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/pkg_workflows.Workflow"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/groups/{group_id}/workflows/{workflow_id}": {
            "get": {
                "description": "return the workflow definition registered.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get a workflow definition",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/pkg_workflows.Workflow"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "httputil.HTTPError": {
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
        },
        "pkg_workflows.Group": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 16,
                    "minLength": 4,
                    "example": "kogito-examples"
                },
                "repository": {
                    "type": "string",
                    "minLength": 10,
                    "example": "https://github.com/kiegroup/kogito-examples/tree/stable"
                }
            }
        },
        "pkg_workflows.GroupDetails": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 16,
                    "minLength": 4,
                    "example": "kogito-examples"
                },
                "repository": {
                    "type": "string",
                    "minLength": 10,
                    "example": "https://github.com/kiegroup/kogito-examples/tree/stable"
                },
                "workflows": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/pkg_workflows.Workflow"
                    }
                }
            }
        },
        "pkg_workflows.Workflow": {
            "type": "object",
            "properties": {
                "host": {
                    "type": "string"
                },
                "input_arguments": {
                    "type": "string",
                    "example": "{ 'fahrenheit': 100 }"
                },
                "meta": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string",
                    "minLength": 3,
                    "example": "fahrenheit_to_celsius"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "parodos-dev:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Parodos API Documentation",
	Description:      "This is a project to run enterprise workflows on demand",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
