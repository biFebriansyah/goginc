package pkg

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func Server(router *gin.Engine) *http.Server {
	var addr string = "0.0.0.0:8080"
	var whiteList []string

	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}

	if origin := os.Getenv("WHITELIST"); origin != "" {
		whiteList = strings.Split(origin, ",")
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   whiteList,
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
