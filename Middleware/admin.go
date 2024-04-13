package Middleware

import (
	"context"
	"log"
	"mongo/DB"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)
func AdminMiddleware(next echo.HandlerFunc)echo.HandlerFunc{
 return func(c echo.Context) error {
  client:= DB.Client
  adminCollection := client.Database("CourseApp").Collection("Admin")
  username := c.Request().Header.Get("username")
  password := c.Request().Header.Get("password")
  filter := bson.M{
	"username":username,
	 "password":password,
  }
  var admin DB.Admin
  err:=adminCollection.FindOne(context.Background(),filter).Decode(&admin)
  if err!=nil{
	if err== mongo.ErrNoDocuments{
	return c.String(http.StatusNotFound,"admin does not exist")
     }
	 log.Fatal(err)
	 return c.JSON(http.StatusInternalServerError,"internal server error")
  }
  return next(c)
 
}
}