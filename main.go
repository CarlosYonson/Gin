package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.New()

	r.GET("/", func(ctx *gin.Context) {
		// ctx.String(200, "Hola, Mundo!")
		ctx.JSON(200, gin.H{
			"message": "Hola, mundo!",
		})
	})

	r.Run(":8080")
}
