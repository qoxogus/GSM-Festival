package rest

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//RunAPI used in main.go
func RunAPI(address string) {
	e := echo.New()
	// e.Use(middleware.Logger())
	tmp := "C:/Users/user/go/src/Gsmfestival-Master-Front"

	e.Static("/", "static")

	e.Use(middleware.Static(tmp))

	e.GET("/", GetMainPage) //서버 검사
	e.GET("/login", Loginpage)
	e.GET("/signup", Signuppage)
	e.GET("/application", Applicationpage)

	e.POST("/loginpage", Signup) //회원정보 insert
	e.POST("/main", Signin)

	e.Logger.Fatal(e.Start(address))
}

//배태현
