package rest

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//RunAPI used in main.go
func RunAPI(address string) {
	tmp := "C:/Users/user/go/src/Gsmfestival-Master"
	e := echo.New()
	// e.Use(middleware.Logger())
	e.Use(middleware.Static(tmp))
	e.GET("/", GetMainPage) //서버 검사
	// e.POST("/mainpage", Mainpage)
	// e.POST("/loginpage", Loginpage)
	e.POST("/signup", Signup)
	e.POST("/signin", Signin)
	e.Logger.Fatal(e.Start(address))
}
