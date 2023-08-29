package handlers

import (
	"biFebriansyah/gogin/config"
	"biFebriansyah/gogin/internal/models"
	"biFebriansyah/gogin/internal/repositories"
	"biFebriansyah/gogin/pkg"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HandlerMovie struct {
	*repositories.RepoMovie
}

func NewMovie(r *repositories.RepoMovie) *HandlerMovie {
	return &HandlerMovie{r}
}

func (h *HandlerMovie) PostData(ctx *gin.Context) {
	movie := models.Movie{}

	if err := ctx.ShouldBind(&movie); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	movie.Movie_banner = ctx.MustGet("image").(string)
	respone, err := h.CreateMovie(&movie)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, respone)

}

func (h *HandlerMovie) PatchData(ctx *gin.Context) {
	movie := models.Movie{}

	if err := ctx.ShouldBind(&movie); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	movie.Movie_banner = ctx.MustGet("image").(string)
	respone, err := h.UpdateMovie(&movie)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, respone)

}

func (h *HandlerMovie) RemoveData(ctx *gin.Context) {
	idMovie := ctx.Param("id")
	data, err := h.RemoveMovie(idMovie)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, data)
}

func (h *HandlerMovie) FetchAllData(ctx *gin.Context) {
	data, err := h.GetAllMovie()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	pkg.NewRes(200, data).Send(ctx)
}

func (h *HandlerMovie) FetchData(ctx *gin.Context) {
	name := ctx.Query("name")
	page := ctx.DefaultQuery("page", "1")
	limit := ctx.DefaultQuery("limit", "10")

	pg, _ := strconv.Atoi(page)
	lm, _ := strconv.Atoi(limit)

	data, err := h.GetMovie(models.Meta{
		Name:  name,
		Page:  pg,
		Limit: lm,
	})

	log.Println(err)

	if err != nil {
		pkg.NewRes(http.StatusBadRequest, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	pkg.NewRes(200, data).Send(ctx)
}
