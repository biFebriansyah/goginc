package pkg

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func Server(router *gin.Engine) *http.Server {
	var addr string = "0.0.0.0:8080"
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "HEAD", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	srv := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 15,
		Handler:      c.Handler(router),
	}

	return srv
}
