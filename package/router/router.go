package router

import (
	"application/package/controllers"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func InitRouter() {
	router = gin.Default()

	web := router.Group("/")
	{
		web.GET("/")
		web.GET("/login")
		web.GET("/logout")
		web.GET("/register")
		web.GET("/forgot")
	}

	v1 := router.Group("/api/v1")
	{
		v1.GET("/login", controllers.GetAllUsers())
		v1.GET("/logout")
		v1.GET("/register")
		v1.GET("/forgot")
	}
}

func Serve(adress string) error {
	return router.Run(adress)
}
