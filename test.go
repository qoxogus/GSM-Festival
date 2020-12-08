package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func signup(c echo.Context) error {
	indextmp := "C:/Users/user/go/src/Gsmfestival-Master/login.html"
	fmt.Println(c.FormValue("classnum"))
	fmt.Println(c.FormValue("myname"))
	fmt.Println(c.FormValue("myemail"))
	fmt.Println(c.FormValue("pwd"))
	fmt.Println(c.FormValue("pwdck"))
	fmt.Println("----------")
	// classnum := "classnum"
	// myname := "myname"
	// myemail := "myemail"
	// pwd := "pwd"
	// pwdck := "pwdck"
	return c.File(indextmp)
}

func login(c echo.Context) error {
	fmt.Println(c.FormValue("email"))
	fmt.Println(c.FormValue("password"))
	return c.File("C:/Users/user/go/src/Gsmfestival-Master/login.html")
}

func main() {
	tmp := "C:/Users/user/go/src/Gsmfestival-Master"
	e := echo.New()
	e.Use(middleware.Static(tmp))
	e.POST("/signup", signup)
	e.POST("/login", login)
	e.Logger.Fatal(e.Start(":1325"))

}
