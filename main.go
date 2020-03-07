package main

import (
	//"log"
	//"net/http"

	//"github.com/gin-contrib/sessions"
	//"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func mainb() {

	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()

	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./public", true)))
	router.Use(gin.Logger())

	router.Run(":4000")
}
