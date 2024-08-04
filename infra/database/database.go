package database

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mongo mongo.Database

func SetupDbConnection() error {

	viper.SetDefault("DATABASE_URI", "mongodb://localhost:27017")
	viper.SetDefault("DATABASE_NAME", "test")

	uri := viper.GetString("DATABASE_URI")
	dbName := viper.GetString("DATABASE_NAME")

	// Create a new client and connect to the server

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return err
	}

	// Send a ping to confirm a successful connection
	var result bson.M
	var database = client.Database(dbName)
	if err := database.RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		return err
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	Mongo = *database

	return nil
}
