package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

var commitVersion string

func main() {
	postPath := os.Args[1]
	fmt.Printf("SimonSays! - Version %s\n", commitVersion)
	r := gin.Default()
	r.GET("/ping", ping)
	r.POST(postPath, postEcho)
	r.Run(":5050")
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func postEcho(c *gin.Context) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	str := buf.String()
	requester := c.ClientIP()
	fmt.Printf("Request From: %s\n", requester)
	fmt.Println(str)
}
