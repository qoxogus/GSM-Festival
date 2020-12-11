package rest

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//RunAPI used in main.go
func RunAPI(address string) {
	tmp := "C:/Users/user/go/src/Gsmfestival-Master-Front"
	e := echo.New()
	// e.Use(middleware.Logger())
	e.Static("/", "static")
	e.Use(middleware.Static(tmp))
	e.GET("/", GetMainPage)      //서버 검사
	e.POST("/loginpage", Signup) //회원정보 insert
	e.POST("/mainpage", Signin)
	// a :=
	// if a != false (err error) {
	// 	if err != nil {
	// 		return err
	// 	}
	// 	u.ID
	// }
	e.Logger.Fatal(e.Start(address))
}

//배태현 공부!!!
