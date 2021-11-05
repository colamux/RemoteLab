package main

import (
	"fmt"
	"rlab/config"
	"rlab/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// server := routers.Init(signup.Routers, login.Routers, home.Routers, sessions.Routers)
	e := gin.Default()

	// e.Use(middleware.AuthMiddleware())

	config.ParseConfig()
	fmt.Println(config.GlobalConfig)

	router.Init(e, router.Admin, router.Api, router.Root)
	e.Static("/var", "../frontend")

	e.Run("localhost:8000")
}
