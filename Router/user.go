package Router

import (
	"context"

	"log"
	"mongo/DB"
	"mongo/Middleware"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

func UserRoute(e *echo.Echo) {
	e.GET("/user/course", getcourse) //fucntion is directly called from admin.go
	e.GET("/user/course/:id", coursePurchase, Middleware.UserMiddleware)
	e.POST("/user/usersignup", userSignUp)
	e.GET("/user/course", getPurchasedCourse, Middleware.UserMiddleware)
	/*   e.GET("/user/middleware",func(c echo.Context) error {
		return c.String(http.StatusOK,"middlware working correctly")
	  },Middleware.UserMiddleware)
	*/
}
func userSignUp(c echo.Context) error {
	user := new(DB.User)
	if err := c.Bind(user); err != nil {
		log.Fatal(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	user.PurchasedCourse = []primitive.ObjectID{}
	collection := client.Database("CourseApp").Collection("User")
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, "user created sucessfully")
}
func coursePurchase(c echo.Context) error {
	courseId := c.Param("id")
	username := c.Get("username").(string)

	userCollection := client.Database("CourseApp").Collection("User")
	var user DB.User

	filter := bson.M{"username": username}
	err := userCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, "Internal server error")
	}

	/*    // Update the user document to add the course ID to purchasedCourse
	if user.PurchasedCourse == nil || reflect.TypeOf(user.PurchasedCourse).Kind() != reflect.Slice {
		user.PurchasedCourse = make([]primitive.ObjectID, 0)
	} */

	update := bson.M{
		"$addToSet": bson.M{"purchasedCourse": courseId},
	}
	// Ensure that the username and purchasedCourse fields are initialized correctly

	_, err = userCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Printf("Failed to update user: %s", err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, "course purchased successfully")
}

func getPurchasedCourse(c echo.Context) error {
	username := c.Get("username").(string)
	collection := client.Database("CourseApp").Collection("User")
	courseCollection := client.Database("CourseApp").Collection("Course")

	filter := bson.D{{"username", username}}
	var user DB.User
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "internal error")
	}
	var courses []string
	for _, id := range user.PurchasedCourse {
		courseFilter := bson.D{{"_id", id}}
		var course DB.Course
		err := courseCollection.FindOne(context.Background(), courseFilter).Decode(&course)
		if err != nil {
			log.Println(err)
			continue
		}
		courses = append(courses, course.Title)
	}

	return c.JSON(http.StatusOK, courses)
}
