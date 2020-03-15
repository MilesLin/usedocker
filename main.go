package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// todo: flags, SSL

	// todo: Authorization
	// todo: enable SSL
	// todo: api port
	// todo: example test

	r := gin.Default()

	r.POST("/rmi", RemoveImageApi)
	r.POST("/pullimage", PullImageApi)
	r.POST("/stopcontainer", StopContainerApi)
	r.POST("/rm", RemoveContainerApi)
	r.POST("/run", RunContainerApi)

	r.POST("/updaterunningcontainer", UpdateRunningContainerApi)

	log.Fatal(r.Run()) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
