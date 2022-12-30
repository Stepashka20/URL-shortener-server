package main

import (
	"os"

	"math/rand"
	"time"

	"stepashka20/url-shortener/routes"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	router := routes.Init()
	router.UrlRoute()
	router.Run("localhost:" + os.Getenv("PORT"))
}
