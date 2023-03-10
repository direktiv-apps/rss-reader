// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Fetch RSS feeds",
    "title": "rss-reader",
    "version": "1.0",
    "x-direktiv-meta": {
      "categories": [
        "social",
        "network"
      ],
      "container": "gcr.io/direktiv/functions/rss-reader",
      "issues": "https://github.com/direktiv-apps/rss-reader/issues",
      "license": "[Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)",
      "long-description": "This function fetches RSS feeds and converts them into JSON. It can limit the results and has a simple text search.",
      "maintainer": "[direktiv.io](https://www.direktiv.io) ",
      "url": "https://github.com/direktiv-apps/rss-reader"
    }
  },
  "paths": {
    "/": {
      "post": {
        "parameters": [
          {
            "type": "string",
            "default": "development",
            "description": "direktiv action id is an UUID. \nFor development it can be set to 'development'\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          },
          {
            "type": "string",
            "default": "/tmp",
            "description": "direktiv temp dir is the working directory for that request\nFor development it can be set to e.g. '/tmp'\n",
            "name": "Direktiv-TempDir",
            "in": "header"
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "limit": {
                  "description": "Limit number of news items and 0 returns all items.",
                  "type": "integer",
                  "example": 10
                },
                "rss": {
                  "description": "URL of the news feed",
                  "type": "string",
                  "example": "https://www.theonion.com/rss"
                },
                "search": {
                  "description": "Simple search term for RSS items",
                  "type": "string",
                  "example": "hello"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "List of executed commands.",
            "schema": {
              "type": "object",
              "additionalProperties": false
            },
            "examples": {
              "rss-reader": {
                "Channel": {
                  "Items": [
                    {
                      "Content": null,
                      "Description": "Caribbean and UN leaders were on hand for the first regional launch of a plan to ensure that every person on the planet is protected by early warning systems by 2027, held in Bridgetown, Barbados on Tuesday.?",
                      "Guid": "https://news.un.org/feed/view/en/story/2023/02/1133247",
                      "Link": "https://news.un.org/feed/view/en/story/2023/02/1133247",
                      "PubDate": "Tue, 07 Feb 2023 12:00:00 +0000",
                      "Source": "UN News - Global perspective Human stories",
                      "Title": "Caribbean sees first regional launch of global plan on early warning systems"
                    }
                  ]
                }
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            },
            "headers": {
              "Direktiv-ErrorCode": {
                "type": "string"
              },
              "Direktiv-ErrorMessage": {
                "type": "string"
              }
            }
          }
        },
        "x-direktiv": {
          "cmds": [
            {
              "action": "exec",
              "exec": "/bin/rss {{ .Rss }} {{ .Limit }} {{ .Search }}"
            }
          ],
          "output": "{\n  \"rss-reader\": {{ index (index . 0) \"result\" | toJson }}\n}\n"
        },
        "x-direktiv-errors": {
          "io.direktiv.command.error": "Command execution failed",
          "io.direktiv.output.error": "Template error for output generation of the service",
          "io.direktiv.ri.error": "Can not create information object from request"
        },
        "x-direktiv-examples": [
          {
            "content": "- id: rss-reader\n  type: action\n  action:\n    function: rss-reader\n    input: \n      rss: https://news.un.org/feed/subscribe/en/news/region/americas/feed/rss.xml",
            "title": "Basic"
          },
          {
            "content": "- id: rss-reader\n  type: action\n  action:\n    function: rss-reader\n    input: \n      rss: https://news.un.org/feed/subscribe/en/news/region/americas/feed/rss.xml\n      limit: 1\n      search: people",
            "title": "Advanced"
          }
        ],
        "x-direktiv-function": "functions:\n- id: rss-reader\n  image: gcr.io/direktiv/functions/rss-reader:1.0\n  type: knative-workflow"
      },
      "delete": {
        "parameters": [
          {
            "type": "string",
            "description": "On cancel Direktiv sends a DELETE request to\nthe action with id in the header\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          }
        ],
        "responses": {
          "200": {
            "description": ""
          }
        },
        "x-direktiv": {
          "cancel": "echo 'cancel {{ .DirektivActionID }}'"
        }
      }
    }
  },
  "definitions": {
    "direktivFile": {
      "type": "object",
      "x-go-type": {
        "import": {
          "package": "github.com/direktiv/apps/go/pkg/apps"
        },
        "type": "DirektivFile"
      }
    },
    "error": {
      "type": "object",
      "required": [
        "errorCode",
        "errorMessage"
      ],
      "properties": {
        "errorCode": {
          "type": "string"
        },
        "errorMessage": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Fetch RSS feeds",
    "title": "rss-reader",
    "version": "1.0",
    "x-direktiv-meta": {
      "categories": [
        "social",
        "network"
      ],
      "container": "gcr.io/direktiv/functions/rss-reader",
      "issues": "https://github.com/direktiv-apps/rss-reader/issues",
      "license": "[Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)",
      "long-description": "This function fetches RSS feeds and converts them into JSON. It can limit the results and has a simple text search.",
      "maintainer": "[direktiv.io](https://www.direktiv.io) ",
      "url": "https://github.com/direktiv-apps/rss-reader"
    }
  },
  "paths": {
    "/": {
      "post": {
        "parameters": [
          {
            "type": "string",
            "default": "development",
            "description": "direktiv action id is an UUID. \nFor development it can be set to 'development'\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          },
          {
            "type": "string",
            "default": "/tmp",
            "description": "direktiv temp dir is the working directory for that request\nFor development it can be set to e.g. '/tmp'\n",
            "name": "Direktiv-TempDir",
            "in": "header"
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/postParamsBody"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "List of executed commands.",
            "schema": {
              "type": "object",
              "additionalProperties": false
            },
            "examples": {
              "rss-reader": {
                "Channel": {
                  "Items": [
                    {
                      "Content": null,
                      "Description": "Caribbean and UN leaders were on hand for the first regional launch of a plan to ensure that every person on the planet is protected by early warning systems by 2027, held in Bridgetown, Barbados on Tuesday.?",
                      "Guid": "https://news.un.org/feed/view/en/story/2023/02/1133247",
                      "Link": "https://news.un.org/feed/view/en/story/2023/02/1133247",
                      "PubDate": "Tue, 07 Feb 2023 12:00:00 +0000",
                      "Source": "UN News - Global perspective Human stories",
                      "Title": "Caribbean sees first regional launch of global plan on early warning systems"
                    }
                  ]
                }
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            },
            "headers": {
              "Direktiv-ErrorCode": {
                "type": "string"
              },
              "Direktiv-ErrorMessage": {
                "type": "string"
              }
            }
          }
        },
        "x-direktiv": {
          "cmds": [
            {
              "action": "exec",
              "exec": "/bin/rss {{ .Rss }} {{ .Limit }} {{ .Search }}"
            }
          ],
          "output": "{\n  \"rss-reader\": {{ index (index . 0) \"result\" | toJson }}\n}\n"
        },
        "x-direktiv-errors": {
          "io.direktiv.command.error": "Command execution failed",
          "io.direktiv.output.error": "Template error for output generation of the service",
          "io.direktiv.ri.error": "Can not create information object from request"
        },
        "x-direktiv-examples": [
          {
            "content": "- id: rss-reader\n  type: action\n  action:\n    function: rss-reader\n    input: \n      rss: https://news.un.org/feed/subscribe/en/news/region/americas/feed/rss.xml",
            "title": "Basic"
          },
          {
            "content": "- id: rss-reader\n  type: action\n  action:\n    function: rss-reader\n    input: \n      rss: https://news.un.org/feed/subscribe/en/news/region/americas/feed/rss.xml\n      limit: 1\n      search: people",
            "title": "Advanced"
          }
        ],
        "x-direktiv-function": "functions:\n- id: rss-reader\n  image: gcr.io/direktiv/functions/rss-reader:1.0\n  type: knative-workflow"
      },
      "delete": {
        "parameters": [
          {
            "type": "string",
            "description": "On cancel Direktiv sends a DELETE request to\nthe action with id in the header\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          }
        ],
        "responses": {
          "200": {
            "description": ""
          }
        },
        "x-direktiv": {
          "cancel": "echo 'cancel {{ .DirektivActionID }}'"
        }
      }
    }
  },
  "definitions": {
    "direktivFile": {
      "type": "object",
      "x-go-type": {
        "import": {
          "package": "github.com/direktiv/apps/go/pkg/apps"
        },
        "type": "DirektivFile"
      }
    },
    "error": {
      "type": "object",
      "required": [
        "errorCode",
        "errorMessage"
      ],
      "properties": {
        "errorCode": {
          "type": "string"
        },
        "errorMessage": {
          "type": "string"
        }
      }
    },
    "postParamsBody": {
      "type": "object",
      "properties": {
        "limit": {
          "description": "Limit number of news items and 0 returns all items.",
          "type": "integer",
          "example": 10
        },
        "rss": {
          "description": "URL of the news feed",
          "type": "string",
          "example": "https://www.theonion.com/rss"
        },
        "search": {
          "description": "Simple search term for RSS items",
          "type": "string",
          "example": "hello"
        }
      },
      "x-go-gen-location": "operations"
    }
  }
}`))
}
