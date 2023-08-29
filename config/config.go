package config

import (
	"context"

	"github.com/gin-contrib/cors"
)

type Operation func(ctx context.Context) error

type Metas struct {
	Next  interface{} `json:"next"`
	Prev  interface{} `json:"prev"`
	Total int         `json:"total"`
}

type Result struct {
	Data    interface{}
	Meta    interface{}
	Message interface{}
}

var CorsConfig = cors.Config{
	AllowOrigins:     []string{"https://glistening-taffy-bfdcf2.netlify.app", "http://127.0.0.1:3000"},
	AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "HEAD", "OPTIONS"},
	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	AllowCredentials: true,
}
