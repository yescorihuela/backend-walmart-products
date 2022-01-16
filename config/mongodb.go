package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// const MONGODB_TEMPLATE_URI = "mongodb://%s:%s@%s:%s/%s?authenticationDatabase=admin&authSource=admin&readPreference=primary&directConnection=true&ssl=false"
const MONGODB_TEMPLATE_URI = "mongodb://%s:%s"

func BuildMongoURI() string {
	mongoPort := os.Getenv("MONGO_INITDB_PORT")
	mongoDbHost := os.Getenv("MONGO_INITDB_HOST")
	return fmt.Sprintf(MONGODB_TEMPLATE_URI, mongoDbHost, mongoPort)
}

func ConnectToMongoDB() *mongo.Client {
	credentials := options.Credential{
		Username: os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
		Password: os.Getenv("MONGO_INITDB_ROOT_PASSWORD"),
	}
	clientOpts := options.Client().ApplyURI(BuildMongoURI()).SetAuth(credentials)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")

	return client
}

func GetCollectionFromMongoDB(client *mongo.Client, collectionName string) *mongo.Collection {
	mongoDbName := os.Getenv("MONGO_INITDB_DB_NAME")
	collection := client.Database(mongoDbName).Collection(collectionName)
	return collection
}
