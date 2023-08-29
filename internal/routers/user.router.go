package routers

import (
	"biFebriansyah/gogin/internal/handlers"
	"biFebriansyah/gogin/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func user(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/users")

	repo := repositories.NewUser(d)
	handler := handlers.NewUser(repo)

	route.POST("/", handler.PostData)
	route.GET("/", handler.FetchAll)
}
