package DB
import (
	
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//mongo password -p3fPlJQI0LZsftj8
type Admin struct{
	Username string `bson:"username"`
	Password string  `bson:"password"`

}
type User struct{
	Username string `bson:"username"`
	Password string `bson:"password"`
	PurchasedCourse []primitive.ObjectID `bson:"purchasedCourse"`
}
type Course struct{
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Title string `bson:"title"`
	Price int    `bson:"price"`
	Description string `bson:"description"`
	ImageLink string `bson:"imagelink"`
}