package main

import (
	"github.com/TobiasYin/water_drop/api"
	"github.com/TobiasYin/water_drop/middlewares"
	_ "github.com/TobiasYin/water_drop/model"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "hello world",
		})
	})
	middlewares.InitMiddleware(r)
	api.InitUrl(r)

	if err := r.Run("0.0.0.0:8080"); err != nil {
		panic(err)
	}
}
