package handlers

import (
	"biFebriansyah/gogin/internal/models"
	"biFebriansyah/gogin/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerMovie struct {
	*repositories.RepoMovie
}

func NewMovie(r *repositories.RepoMovie) *HandlerMovie {
	return &HandlerMovie{r}
}

func (h *HandlerMovie) PostData(ctx *gin.Context) {
	var movie models.Movie

	if err := ctx.ShouldBind(&movie); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, err := h.CreateMovie(&movie)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, response)
}
