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

func UserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Request().Header.Get("username")
		password := c.Request().Header.Get("password")
		client := DB.Client
		userCollection := client.Database("CourseApp").Collection("User")
		filter := bson.M{
			"username": username,
			"password": password,
		}
		var user DB.User
		err := userCollection.FindOne(context.Background(), filter).Decode(&user)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return c.String(http.StatusNotFound, "user does not exist")
			}
			log.Println(err)
			return c.String(http.StatusInternalServerError, "internal server error in decoding in middleware")
		}
		c.Set("username", username)
		return next(c)
	}
}
