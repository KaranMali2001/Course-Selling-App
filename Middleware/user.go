package Middleware

import (
//	"context"
	"errors"
	"log"
	//"mongo/DB"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/mongo"
)

/* func UserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
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
} */
func VerifyToken(tokenString string)(string,error){
	token,err:=jwt.Parse(tokenString,func(t *jwt.Token) (interface{}, error) {
		return ([]byte(JWT_SECRAT)),nil
	})
	if err!=nil{
		log.Println(err)
		return "",err
	}
	claims,ok:= token.Claims.(jwt.MapClaims)
	if !ok{
		return "",errors.New("invalid claims ")
	}
	username,ok:=claims["username"].(string)
	if !ok{
		return "",errors.New("username not found in string")
	
	 }
	 return username,nil
}
//middleware with jwt
func UserMiddleware(next echo.HandlerFunc)echo.HandlerFunc{
	return func(c echo.Context) error {
		token:= c.Request().Header.Get("authorization")
		if token==""{
			return c.String(http.StatusBadRequest,"empty token")
		}
		tokenString := strings.TrimPrefix(token,"bearer ")
		tokenValidation,err:=jwt.Parse(tokenString,func(t *jwt.Token) (interface{}, error) {
			return ([]byte(JWT_SECRAT)),nil
		})
		if err!=nil || !tokenValidation.Valid{
			log.Println(err)
			return c.String(http.StatusInternalServerError,"internal server error")
		}
		username,err:=VerifyToken(tokenString)
		if err!=nil{
             log.Println(err)
			 return c.String(http.StatusInternalServerError,"server error")
		}
		c.Set("username",username)
		return next(c)
	}
}
