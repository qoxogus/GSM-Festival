package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func signup(c echo.Context) error {
	fmt.Println(c.FormValue("classnum"))
	fmt.Println(c.FormValue("myname"))
	fmt.Println(c.FormValue("myemail"))
	fmt.Println(c.FormValue("pwd"))
	fmt.Println(c.FormValue("pwdck"))
	return c.File("C:/Users/user/go/src/Gsmfestival-Master/signup.html")
}

func main() {
	tmp := "C:/Users/user/go/src/Gsmfestival-Master"
	e := echo.New()
	e.Use(middleware.Static(tmp))
	e.POST("/signup", signup)
	e.Logger.Fatal(e.Start(":1324"))

}
