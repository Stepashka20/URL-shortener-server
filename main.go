package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"crypto/rand"
	"encoding/base64"
)

// ShortURL represents a shortened URL in the database
type ShortURL struct {
	Key    string `bson:"key"`
	URL    string `bson:"url"`
	Clicks int    `bson:"clicks"`
}

var client *mongo.Client = ConnectToDB()

func main() {
	router := gin.Default()
	router.GET("/:key", redirect)
	router.POST("/", create)

	// Start the server
	router.Run(":7999")
}

func ConnectToDB() *mongo.Client {
	// connect env file
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://Stepashka20:sck3jv8mpkocnbvcm4@127.0.0.1:27017/"))

	if err != nil {
		log.Fatal(err)
	}

	// Create connect
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}

func redirect(c *gin.Context) {
	// Get the key from the URL path
	key := c.Param("key")

	// Find the corresponding URL in the database
	var shortURL ShortURL
	err := findOne(bson.M{"key": key}, &shortURL)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// Redirect to the full URL
	c.Redirect(http.StatusMovedPermanently, shortURL.URL)
}
func generateKey(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}
func create(c *gin.Context) {
	// Parse the form data
	var form struct {
		URL string `form:"url" binding:"required"`
	}
	if err := c.Bind(&form); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	// print form
	log.Println(form)

	key, _ := generateKey(8)

	shortURL := ShortURL{
		Key:    key,
		URL:    form.URL,
		Clicks: 0,
	}

	_, err := client.Database("urls").Collection("shorturls").InsertOne(context.Background(), shortURL)
	if err != nil {
		log.Fatal(err)
	}

	c.String(http.StatusOK, key)
}

func findOne(filter interface{}, result interface{}) error {
	// Use the "shorturls" collection in the "test" database
	collection := client.Database("urls").Collection("shorturls")

	// Find the document
	err := collection.FindOne(context.TODO(), filter).Decode(result)
	if err == mongo.ErrNoDocuments {
		return errors.New("document not found")
	}
	if err != nil {
		return err
	}
	return nil
}
