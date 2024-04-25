package DB

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var clientOption *options.ClientOptions
var Client *mongo.Client

func init() {
	clientOption = options.Client().ApplyURI("mongodb://localhost:27017/")
	var err error
	Client, err = mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	err = Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongo connected sucessfully")
}
