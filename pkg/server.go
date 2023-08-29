package pkg

import (
	"biFebriansyah/gogin/config"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Server(router *gin.Engine) *http.Server {
	var addr string = "0.0.0.0:8080"
	var whiteList []string
	var corsConfig = config.CorsConfig

	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}

	if origin := os.Getenv("WHITELIST"); origin != "" {
		whiteList = strings.Split(origin, ",")
	}

	log.Println(whiteList)

	corsConfig.AllowOriginFunc = func(origin string) bool {
		var check bool
		for _, v := range whiteList {
			log.Println(v)
			if v == origin {
				check = true
				break
			}
		}

		log.Printf("origin %s allowing %t", origin, check)
		return check
	}

	router.Use(cors.New(corsConfig))
	srv := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 15,
		Handler:      router,
	}

	return srv
}
