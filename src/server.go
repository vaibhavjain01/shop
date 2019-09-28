package main

import (
	"net/http"
	"fmt"
	
	"github.com/labstack/echo"
)

func StartServer() {
	e := echo.New()
	//e.POST("/normal", normal)
	e.GET("/pathparam/:id", pathparam)
	e.GET("/normal", normal)
	e.GET("/queryparam", queryparam)
	//e.PUT("/users/:id", updateUser)
	//e.DELETE("/users/:id", deleteUser)
	e.Logger.Fatal(e.Start(":1323"))
}

func normal(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World Nomral")
}

func pathparam(c echo.Context) error {
	id := c.Param("id")
	fmt.Println(id)
	return c.String(http.StatusOK, "Hello, World! Path Param")
}

func queryparam(c echo.Context) error {
	id := c.QueryParam("id")
	fmt.Println(id)
	return c.String(http.StatusOK, "Hello, World! Path Param")
}