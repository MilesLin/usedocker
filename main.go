package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	//_ "github.com/swaggo/gin-swagger/example/basic/docs" // docs is generated by Swag CLI, you have to import it.
	_ "github.com/mileslin/usedocker/docs"
	"log"
)

func main() {
	// todo: document api
	// todo: document readme
	// https://github.com/swaggo/swag
	var enableSSL = flag.Bool("enableSSL", false, "To enable SSL by adding -enableSSL flag")
	var sslport = flag.String("sslport", "443", "The port of https. 443 is default value.")
	var port = flag.String("port", "8080", "The port of http. 8080 is default value.")

	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Authorization
	var acct = flag.String("account", "", "The account for basic authorization usage.")
	var pwd = flag.String("password", "", "The password for basic authorization usage.")
	var swag = flag.Bool("swag", false, "To enable swagger by adding -swag flag. The swagger page is /swagger/index.html")

	flag.Parse()

	r := gin.Default()
	if *swag {
		url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}

	if *acct != "" && *pwd != "" {
		r.Use(gin.BasicAuth(gin.Accounts{
			*acct: *pwd,
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
