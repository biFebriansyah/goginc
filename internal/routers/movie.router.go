package routers

import (
	"biFebriansyah/gogin/internal/handlers"
	"biFebriansyah/gogin/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// ! /movie
func movie(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/movie")

	// dependcy injection
	repo := repositories.NewMovie(d)
	handler := handlers.NewMovie(repo)

	route.POST("/", handler.PostData)

}
