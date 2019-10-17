package main

import (
	"github.com/labstack/echo"
	"github.com/yrisob/cms_core/handlers"
)

func initRouter(e *echo.Echo) {
	//	e.GET("/test", handlers.GetUsers)
	//	e.GET("/test/:id", handlers.GetUser)
	e.POST("/test", handlers.PostTest)
}
