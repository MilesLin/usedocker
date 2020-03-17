package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// todo: example test
	// todo: refactor
	// todo: enable swagger https://github.com/swaggo/gin-swagger
	var enableSSL = flag.Bool("enableSSL", false, "To enable SSL by adding -enableSSL flag")
	var sslport = flag.String("sslport", "443", "The port of https. 443 is default value.")
	var port = flag.String("port", "8080", "The port of http. 8080 is default value.")

	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Authorization
	var acct = flag.String("account", "", "The account for basic authorization usage.")
	var pwd = flag.String("password", "", "The password for basic authorization usage.")

	flag.Parse()

	r := gin.Default()

	if *acct != "" && *pwd != "" {
		r.Use(gin.BasicAuth(gin.Accounts{
			*acct:*pwd,
		}))
	}

	r.POST("/rmi", RemoveImageApi)
	r.POST("/pull", PullImageApi)
	r.POST("/stop", StopContainerApi)
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
