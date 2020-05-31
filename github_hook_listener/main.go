package main

import (
	"github.com/TobiasYin/water_drop/github_hook_listener/event"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.New()

	r.POST("/", func(context *gin.Context) {
		var e event.Event
		if err := context.ShouldBind(&e); err != nil {
			context.JSON(400, gin.H{"Error": err})
			return
		}
		if event.IsTarget(e) {
			log.Println("Found Target Event!")
			event.Workflow()
		}
	})

	r.GET("/restart", func(context *gin.Context) {
		event.Workflow()
	})

	if err := r.Run("0.0.0.0:9090"); err != nil {
		panic(err)
	}
}
