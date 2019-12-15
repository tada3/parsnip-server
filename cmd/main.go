package main

import (
	"fmt"

	"flag"
	"git-dev.linecorp.com/tadafumi-yoshihara/idpgo/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("idpgo started!")

	flag.Parse()
	args := flag.Args()
	fmt.Println(args)

	if len(args) > 0 {
		env := args[0]
		handler.SetEnv(env)
	}

	router := gin.Default()
	router.LoadHTMLGlob("dist/*.html")

	router.GET("/", handler.Root)

	router.POST("/logout", handler.Logout)
	router.GET("/logout", handler.Logout)
	router.GET("/login", handler.Login)

	router.GET("api/v1/getLogin", handler.GetLoginStatus)
	router.GET("api/v1/logoutUrl", handler.GetLogoutURL)

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Static("/js", "./dist/js")

	router.Run()

	fmt.Println("done")
}
