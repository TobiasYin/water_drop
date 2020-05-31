package main

import (
	_ "github.com/TobiasYin/water_drop/model"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "hello world test new",
		})
	})

	if err := r.Run("0.0.0.0:8080"); err != nil {
		panic(err)
	}
}
