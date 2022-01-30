package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the Gin framework
	r := gin.Default()

	// Tmp example endpoint
	r.GET("/tmp", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "tmp hello!",
		})
	})

	// Run the API, and set the port
	if err := r.Run(":80"); err != nil {
		fmt.Print("Unable to start the API :(")
	}
}
