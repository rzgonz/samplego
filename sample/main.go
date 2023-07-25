package main

import (
	"fmt"
	"sample/api/response"
	"sample/helpers"
	"sample/helpers/string"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello again")
	helpers.TestHelper()
	string.PrintDummy()
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/list", func(ctx *gin.Context) {
		var albums = []response.SampleResponse{
			{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
			{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
			{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
		}
		ctx.JSON(200, gin.H{
			"data": albums,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
