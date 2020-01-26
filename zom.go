package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yuliandong/zoman/imgana"
)

func main() {
	fmt.Println("New zom man!")
	imgana.Example()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(400, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
