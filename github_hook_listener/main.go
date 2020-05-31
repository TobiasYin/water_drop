package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "hello",
		})
	})

	if err := r.Run(); err != nil {
		panic(err)
	}
}
