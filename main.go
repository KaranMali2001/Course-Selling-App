package main

import (
	
	"net/http"
    _"mongo/DB"
	"mongo/Router"
	// "mongo/middleware"
	"github.com/labstack/echo/v4"
	
)


//creating new schema in mongo db


func greating(c echo.Context)error{
 return c.JSON(http.StatusOK,"server started ")
}
func main(){
	e :=echo.New()
	e.HideBanner=true
	e.GET("/",greating)
    Router.AdminRoute(e)
	Router.UserRoute(e)
	e.Start(":8080")
}