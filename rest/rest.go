package rest

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//RunAPI used in main.go
func RunAPI(address string) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", GetMainPage) //서버 검사
	e.POST("/signup", Signup)
	e.POST("/signin", Signin)
	e.Logger.Fatal(e.Start(address))
}
