package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SayHello(name string) string {
	return "hello " + name
}

func Example() {
	router := gin.Default()

	router.GET("/", examples)
	router.GET("/query", queryString)
	router.GET("/params/:user/:slug", paramsString)
	router.POST("/login", reqBody)

	router.Run(":8080")
}

func examples(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status":  200,
		"message": "Hello from gin",
	})
}

type qString struct {
	Limit string `form:"limit"`
	Page  string `form:"page"`
}

// ! localhost:8080/query?page=1&limit10
func queryString(ctx *gin.Context) {
	// page := ctx.Query("page")
	// limit := ctx.Query("limit")

	var data qString
	if err := ctx.ShouldBind(&data); err != nil {
		log.Println(err)
	}

	ctx.JSON(200, gin.H{
		"page":  data.Page,
		"limit": data.Limit,
	})
}

type pString struct {
	User string `uri:"user"`
	Slug string `uri:"slug"`
}

// ! localhost:8080/params/:username/:slug
func paramsString(ctx *gin.Context) {
	// username := ctx.Param("user")
	// slug := ctx.Param("slug")

	var data pString
	if err := ctx.ShouldBindUri(&data); err != nil {
		log.Println(err)
	}

	ctx.JSON(200, gin.H{
		"username": data.User,
		"slug":     data.Slug,
	})
}

type body struct {
	Username string `form:"username"`
	Password string `form:"username"`
}

func reqBody(ctx *gin.Context) {

	file, err := ctx.FormFile("image")
	if err != nil {
		log.Println(err)
	}

	log.Println(file.Filename)

	var data body
	if err := ctx.ShouldBind(&data); err != nil {
		log.Println(err)
	}

	ctx.JSON(200, data)

}
