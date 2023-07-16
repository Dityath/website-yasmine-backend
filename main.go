package main

import (
	"github.com/dityath/website-yasmine-backend/initializers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	r.Run()
}
