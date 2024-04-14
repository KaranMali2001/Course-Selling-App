package Router

import (
	"context"
	"log"
	"mongo/DB"

	"mongo/Middleware"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func init() {

	client = DB.Client
}
func AdminRoute(e *echo.Echo) {
	e.POST("/admin/signup", adminsignup)
	e.POST("/admin/createcourse", createcourse, Middleware.AdminMiddleware)
	e.GET("/admin/course", getcourse, Middleware.AdminMiddleware)
}

func adminsignup(c echo.Context) error {
	admin := new(DB.Admin)
	if err := c.Bind(admin); err != nil {
		log.Fatal(err)
		return c.String(http.StatusBadRequest, "invalid input")
	}
	collection := client.Database("CourseApp").Collection("Admin")
	_, err := collection.InsertOne(context.TODO(), admin)
	if err != nil {
		log.Fatal(err)
		return c.String(http.StatusInternalServerError, "internal server err")
	}
	return c.String(http.StatusOK, "admin account created sucessfully ")

}
func createcourse(c echo.Context) error {
	course := new(DB.Course)
	if err := c.Bind(course); err != nil {
		log.Fatal(err)
		return c.String(http.StatusBadRequest, "invalid input")
	}
	collection := client.Database("CourseApp").Collection("Course")
	result, err := collection.InsertOne(context.TODO(), course)
	if err != nil {
		log.Fatal(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	id := result.InsertedID
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "course created sucessfully",
		"course_id": id,
	})
}
func getcourse(c echo.Context) error {
	collection := client.Database("CourseApp").Collection("Course")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	defer cursor.Close(context.TODO())
	var courses []DB.Course
	for cursor.Next(context.TODO()) {
		var course DB.Course
		if err := cursor.Decode(&course); err != nil {
			log.Fatal(err)
			return c.String(http.StatusInternalServerError, "internal server error")
		}
		courses = append(courses, course)
	}

	return c.JSON(http.StatusOK, courses)
}
