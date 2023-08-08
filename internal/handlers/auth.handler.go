package handlers

import (
	"biFebriansyah/gogin/config"
	"biFebriansyah/gogin/internal/repositories"
	"biFebriansyah/gogin/pkg"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `db:"username" json:"username" form:"username"`
	Password string `db:"password" json:"password,omitempty"`
}

type HandlerAuth struct {
	*repositories.RepoUser
}

func NewAuth(r *repositories.RepoUser) *HandlerAuth {
	return &HandlerAuth{r}
}

func (h *HandlerAuth) Login(ctx *gin.Context) {

	var data User
	if ers := ctx.ShouldBind(&data); ers != nil {
		pkg.NewRes(500, &config.Result{
			Data: ers.Error(),
		}).Send(ctx)
		return
	}

	pkg.NewRes(200, &config.Result{Data: data.Username}).Send(ctx)
}
