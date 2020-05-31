package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	r := gin.New()

	r.POST("/", func(context *gin.Context) {
		data, _ := ioutil.ReadAll(context.Request.Body)
		err := ioutil.WriteFile(fmt.Sprintf("req-%s.txt", time.Now().String()), data, os.ModePerm)
		context.JSON(200, gin.H{
			"info": err,
		})
	})

	if err := r.Run("0.0.0.0:9090"); err != nil {
		panic(err)
	}
}
