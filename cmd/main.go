package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tada3/parsnip-server/handler"
)

func main() {
	fmt.Println("parsni-server started!")

	router := gin.Default()
	router.Use(cors())

	router.GET("/tasks", handler.GetTasks)

	router.POST("/tasks", handler.AddTask)

	router.PUT("/tasks/:taskID", handler.EditTask)

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run()

	fmt.Println("done")
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,HEAD,OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Next()

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
