package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const msgFormat = `
# HELP OnlineUser Number of online player that currently
# TYPE OnlineUser Counter 
OnlineUser %d
`

func main() {
	app := gin.Default()

	app.GET("/metrics", func(c *gin.Context) {
		c.String(200, fmt.Sprintf(msgFormat, 30))
	})

	app.Run(":5000")
}
