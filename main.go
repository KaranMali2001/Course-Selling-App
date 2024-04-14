package main

import (
	"mongo/Router"
	"net/http"
	
	
	"github.com/labstack/echo/v4"
)

//creating new schema in mongo db
var JWT_SECRAT ="Karan_SERVER"
func greating(c echo.Context) error {
	return c.JSON(http.StatusOK, "server started at 8080  ")
}
func main() {
	e := echo.New()
	e.HideBanner = true
	e.GET("/", greating)
	Router.AdminRoute(e)
	Router.UserRoute(e)
	e.Start(":8080")
	
}
