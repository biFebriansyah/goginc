package routers

import (
	"biFebriansyah/gogin/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func New(db *sqlx.DB) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(config.CorsConfig))

	movie(router, db)
	user(router, db)
	auth(router, db)

	return router
}
