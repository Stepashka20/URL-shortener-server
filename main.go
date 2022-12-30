package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"math/rand"
	"time"

	"github.com/joho/godotenv"
)

type ShortURL struct {
	Key    string `bson:"key"`
	URL    string `bson:"url"`
	Clicks int    `bson:"clicks"`
}

var client *mongo.Client = ConnectToDB()

func main() {
	rand.Seed(time.Now().UnixNano())
	router := gin.Default()
	router.GET("/:key", redirect)
	router.POST("/", create)

	router.Run(":7999")
}

func ConnectToDB() *mongo.Client {
	// connect env file
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

	fmt.Println("Connected to MongoDB!")

	return client
}

func redirect(c *gin.Context) {
	key := c.Param("key")

	var shortURL ShortURL
	err := findOne(bson.M{"key": key}, &shortURL)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.Redirect(http.StatusMovedPermanently, shortURL.URL)
}

func generateKey(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
func create(c *gin.Context) {
	var form struct {
		URL string `form:"url" binding:"required"`
	}
	if err := c.Bind(&form); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	key := generateKey(8)

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
