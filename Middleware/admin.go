package Middleware

import (
	//"context"
	//"context"
	
	"errors"
	"log"

	//"mongo/DB"
	//"mongo"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)
var JWT_SECRAT="karan_server"
//without JWT
/* func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		client := DB.Client
		adminCollection := client.Database("CourseApp").Collection("Admin")
		username := c.Request().Header.Get("username")
		password := c.Request().Header.Get("password")
		filter := bson.M{
			"username": username,
			"password": password,
		}
		var admin DB.Admin
		err := adminCollection.FindOne(context.Background(), filter).Decode(&admin)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return c.String(http.StatusNotFound, "admin does not exist")
			}
			log.Fatal(err)
			return c.JSON(http.StatusInternalServerError, "internal server error")
		}
		return next(c)

	}
} */
//with jwt
//verifying jwt
func VerifyJwt(tokenstring string)(string,error){
 token,err:= jwt.Parse(tokenstring,func(t *jwt.Token) (interface{}, error) {
	return []byte(JWT_SECRAT),nil
 })
 if err!=nil || !token.Valid{
	return "",err
 }
 claims,ok:=token.Claims.(jwt.MapClaims)
 if !ok{
	return "",errors.New("invalid token claim")
 }
 username,ok:=claims["username"].(string)
 if !ok{
	return "",errors.New("username not found in string")

 }
 return username,nil
}
func AdminMiddleware(next echo.HandlerFunc)echo.HandlerFunc{
	return func(c echo.Context)error{
		token:= c.Request().Header.Get("authorization")
		if token == "" {
			return c.JSON(http.StatusUnauthorized, "token is not valid")
		}
  		if !strings.HasPrefix(token, "Bearer "){
         		return c.String(http.StatusBadRequest,"missing bearer in header")

		}
		tokenString := strings.TrimPrefix(token,"bearer")
		tokenValidation,err:= jwt.Parse(tokenString,func(t *jwt.Token) (interface{}, error) {
			return []byte(JWT_SECRAT),nil
		})
		if err!=nil || !tokenValidation.Valid{
			log.Println(err)
			return c.JSON(http.StatusUnauthorized,"invalid token")
		}
		
		username,err:=VerifyJwt(tokenString)
		if err!=nil{
            return c.JSON(http.StatusInternalServerError,"not able to verify token")
		}
		c.Set("username",username)
	
		return next(c)
	}
}