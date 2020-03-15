package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// todo: flags, SSL
	// RestartPolicy
	//   Empty string means not to restart
	//   always Always restart
	//   unless-stopped Restart always except when the user has manually stopped the container
	//   on-failure Restart only when the container exit code is non-zero
	// Authorization, port mapping, container name, imageName

	// todo: enable SSL

	// todo: api port

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/rmi", RemoveImageApi)

	log.Fatal(r.Run()) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
