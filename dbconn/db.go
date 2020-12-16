package dbconn

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//DB
func GetDBCollection() (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB (연결)
	fmt.Println("MongoDB Connect")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection(연결검증)
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	// GPM.users DataBase     MongoDB
	collection := client.Database("GPM").Collection("users")

	return collection, nil
}
