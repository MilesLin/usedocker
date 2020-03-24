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
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/pull": {
            "post": {
                "description": "Pull an image with authentication by image name. Your should set -cracct and -crpwd flag for username and password when running the console.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Pull an image with authentication",
                "parameters": [
                    {
                        "description": "the body content",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Image"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "the sample of body is {\\\"msg\\\": \\\"message\\\", \\\"err\\\":\\\"message\\\"}",
                        "schema": {
                            "type": "body"
                        }
                    }
                }
            }
        },
        "/rm": {
            "post": {
                "description": "Remove a container by container name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Remove a container",
                "parameters": [
                    {
                        "description": "the body content",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Container"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "the sample of body is {\\\"msg\\\": \\\"message\\\", \\\"err\\\":\\\"message\\\"}",
                        "schema": {
                            "type": "body"
                        }
                    }
                }
            }
        },
        "/rmi": {
            "post": {
                "description": "Remove an image by image name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Remove an image",
                "parameters": [
                    {
                        "description": "the body content",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Image"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "the sample of body is {\\\"msg\\\": \\\"message\\\", \\\"err\\\":\\\"message\\\"}",
                        "schema": {
                            "type": "body"
                        }
                    }
                }
            }
        },
        "/run": {
            "post": {
                "description": "It do create and start to run the container a container by container name.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Run a container",
                "parameters": [
                    {
                        "description": "the body content",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.ContainerConfig"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "the sample of body is {\\\"msg\\\": \\\"message\\\", \\\"err\\\":\\\"message\\\"}",
                        "schema": {
                            "type": "body"
                        }
                    }
                }
            }
        },
        "/stop": {
            "post": {
                "description": "Stop a container by container name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Stop a container",
                "parameters": [
                    {
                        "description": "the body content",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Container"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "the sample of body is {\\\"msg\\\": \\\"message\\\", \\\"err\\\":\\\"message\\\"}",
                        "schema": {
                            "type": "body"
                        }
                    }
                }
            }
        },
        "/updaterunningcontainer": {
            "post": {
                "description": "It do 1. stop container, 2. remove container 3. remove image 4. pull image 5. run container.  If one step failed, then it stop immediately.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update a running container",
                "parameters": [
                    {
                        "description": "the body content",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.ContainerConfigWithAuth"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "the sample of body is {\\\"msg\\\": \\\"message\\\", \\\"err\\\":\\\"message\\\"}",
                        "schema": {
                            "type": "body"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Container": {
            "type": "object",
            "required": [
                "containerName"
            ],
            "properties": {
                "containerName": {
                    "type": "string",
                    "example": "dockerapp"
                }
            }
        },
        "main.ContainerConfig": {
            "type": "object",
            "required": [
                "containerName",
                "exportPort",
                "hostIP",
                "hostPort",
                "imageNameTag"
            ],
            "properties": {
                "containerName": {
                    "type": "string",
                    "example": "dockerapp"
                },
                "env": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "abc=123",
                        "xyz=999"
                    ]
                },
                "exportPort": {
                    "type": "string",
                    "example": "80"
                },
                "hostIP": {
                    "type": "string",
                    "example": "0.0.0.0"
                },
                "hostPort": {
                    "type": "string",
                    "example": "8080"
                },
                "imageNameTag": {
                    "type": "string",
                    "example": "mileslin/dockerlab:latest"
                },
                "mount": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.Mount"
                    }
                },
                "restartPolicy": {
                    "description": "It supports ` + "`" + `no` + "`" + `, ` + "`" + `always` + "`" + `, ` + "`" + `on-failure` + "`" + `, ` + "`" + `unless-stopped` + "`" + `",
                    "type": "string",
                    "example": "always"
                }
            }
        },
        "main.ContainerConfigWithAuth": {
            "type": "object",
            "required": [
                "containerName",
                "exportPort",
                "hostIP",
                "hostPort",
                "imageNameTag"
            ],
            "properties": {
                "containerName": {
                    "type": "string",
                    "example": "dockerapp"
                },
                "env": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "abc=123",
                        "xyz=999"
                    ]
                },
                "exportPort": {
                    "type": "string",
                    "example": "80"
                },
                "hostIP": {
                    "type": "string",
                    "example": "0.0.0.0"
                },
                "hostPort": {
                    "type": "string",
                    "example": "8080"
                },
                "imageNameTag": {
                    "type": "string",
                    "example": "mileslin/dockerlab:latest"
                },
                "mount": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.Mount"
                    }
                },
                "restartPolicy": {
                    "description": "It supports ` + "`" + `no` + "`" + `, ` + "`" + `always` + "`" + `, ` + "`" + `on-failure` + "`" + `, ` + "`" + `unless-stopped` + "`" + `",
                    "type": "string",
                    "example": "always"
                },
                "withAuth": {
                    "type": "boolean"
                }
            }
        },
        "main.Image": {
            "type": "object",
            "required": [
                "imageNameTag"
            ],
            "properties": {
                "imageNameTag": {
                    "type": "string",
                    "example": "mileslin/dockerlab:latest"
                }
            }
        },
        "main.Mount": {
            "type": "object",
            "properties": {
                "source": {
                    "type": "string",
                    "example": "myvolume"
                },
                "target": {
                    "type": "string",
                    "example": "/app/appdata"
                },
                "type": {
                    "description": "It supports ` + "`" + `bind` + "`" + `, ` + "`" + `volume` + "`" + `, ` + "`" + `tmpfs` + "`" + `, ` + "`" + `npipe` + "`" + `",
                    "type": "string",
                    "example": "volume"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
