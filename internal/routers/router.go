package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func New(db *sqlx.DB) *gin.Engine {
	router := gin.Default()

	movie(router, db)
	user(router, db)
	auth(router, db)

	return router
}
