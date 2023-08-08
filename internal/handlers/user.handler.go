package handlers

import (
	"biFebriansyah/gogin/config"
	"biFebriansyah/gogin/internal/models"
	"biFebriansyah/gogin/internal/repositories"
	"biFebriansyah/gogin/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
	*repositories.RepoUser
}

func NewUser(r *repositories.RepoUser) *HandlerUser {
	return &HandlerUser{r}
}

func (h *HandlerUser) PostData(ctx *gin.Context) {
	var ers error
	data := models.User{
		Role: "user",
	}

	if ers = ctx.ShouldBind(&data); ers != nil {
		ctx.AbortWithError(http.StatusBadRequest, ers)
		return
	}

	// TODO payload validtion here

	data.Password, ers = pkg.HashPassword(data.Password)
	if ers != nil {
		pkg.NewRes(401, &config.Result{
			Data: ers.Error(),
		}).Send(ctx)
		return
	}

	respone, ers := h.CreateUser(&data)
	if ers != nil {
		ctx.AbortWithError(http.StatusBadRequest, ers)
		return
	}

	pkg.NewRes(200, respone).Send(ctx)

}

func (h *HandlerUser) FetchAll(ctx *gin.Context) {
	data, err := h.GetAllUser()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	pkg.NewRes(200, data).Send(ctx)
}