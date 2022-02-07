package helpers

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



func GetDBClient() *mongo.Client {
	// Set client options
	fmt.Print("Called Client")
	clientOptions := options.Client().ApplyURI("mongodb+srv://Sunkanmi:sunkanmi123@cluster0.bpei9.mongodb.net/helpful?retryWrites=true&w=majority")
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	return client
}
