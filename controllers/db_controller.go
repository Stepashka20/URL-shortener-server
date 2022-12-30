package controllers

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client = ConnectToDB()

func ConnectToDB() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URL")))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Connected to MongoDB!")

	return client
}

func findOne(filter interface{}, result interface{}) error {
	collection := client.Database("urls").Collection("shorturls")

	err := collection.FindOne(context.TODO(), filter).Decode(result)
	if err == mongo.ErrNoDocuments {
		return errors.New("document not found")
	}
	if err != nil {
		return err
	}
	return nil
}

func InsertNewUrl(key string, url string) {
	collection := client.Database("urls").Collection("shorturls")

	_, err := collection.InsertOne(context.TODO(), bson.M{"key": key, "url": url, "clicks": 0})
	if err != nil {
		log.Fatal(err)
	}
}
