package docs

import (
    "github.com/swaggo/swag"
)

func init() {
    swag.Register(swag.Name, &s{})
}

type s struct{}

func (s *s) ReadDoc() string {
    return `{
        "swagger": "2.0",
        "info": {
            "description": "Music API with authentication and CRUD functionality",
            "title": "Music API",
            "version": "1.0.0"
        },
        "host": "localhost:8080",
        "basePath": "/api/music",
        "schemes": [
            "http"
        ],
        "paths": {
            "/": {
                "get": {
                    "summary": "Get all musics",
                    "responses": {
                        "200": {
                            "description": "OK"
                        }
                    }
                }
            }
        }
    }`
}
