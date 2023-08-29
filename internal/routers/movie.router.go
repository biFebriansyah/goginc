package routers

import (
	"biFebriansyah/gogin/internal/handlers"
	"biFebriansyah/gogin/internal/middleware"
	"biFebriansyah/gogin/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func movie(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/movie")

	repo := repositories.NewMovie(d)
	handler := handlers.NewMovie(repo)

	route.POST("", middleware.AuthJwt("admin"), middleware.UploadFile, handler.PostData)
	route.PATCH("", middleware.UploadFile, handler.PatchData)
	route.DELETE(":id", handler.RemoveData)
	route.GET("", handler.FetchData)
	route.GET("all", handler.FetchAllData)
}
