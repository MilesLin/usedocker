package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// todo: Authorization
	// todo: example test
	// todo: refactor
	var enableSSL = flag.Bool("enableSSL", false, "To enable SSL by adding -enableSSL flag")
	var sslport = flag.String("sslport", "443", "The port of https. 443 is default value.")
	var port = flag.String("port", "8080", "The port of http. 8080 is default value.")
	flag.Parse()

	r := gin.Default()

	r.POST("/rmi", RemoveImageApi)
	r.POST("/pullimage", PullImageApi)
	r.POST("/stopcontainer", StopContainerApi)
	r.POST("/rm", RemoveContainerApi)
	r.POST("/run", RunContainerApi)

	r.POST("/updaterunningcontainer", UpdateRunningContainerApi)
	if *enableSSL {
		go func() {
			r.RunTLS(":"+*sslport, "./cert.pem", "./key.pem")
		}()
	}

	log.Fatal(r.Run(":" + *port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
