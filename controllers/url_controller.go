package controllers

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type ShortURL struct {
	Key    string `bson:"key"`
	URL    string `bson:"url"`
	Clicks int    `bson:"clicks"`
}

func GetRedirect(c *gin.Context) {
	key := c.Param("key")

	var shortURL ShortURL
	err := getOriginalUrl(bson.M{"key": key}, &shortURL)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.Redirect(http.StatusMovedPermanently, shortURL.URL)
}

func GetShortUrl(c *gin.Context) {
	var form struct {
		URL string `form:"url" binding:"required"`
	}
	if err := c.Bind(&form); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	key := generateKey(8)

	InsertNewUrl(key, form.URL)

	c.String(http.StatusOK, key)
}

func generateKey(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
