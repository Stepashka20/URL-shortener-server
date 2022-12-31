package routes

import (
	"stepashka20/url-shortener/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type routerEngine struct {
	*gin.Engine
}

func Init() *routerEngine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "GET", "OPTIONS"},
	}))
	return &routerEngine{r}
}

func (router *routerEngine) UrlRoute() {
	router.GET("/:key", controllers.GetRedirect)
	router.POST("/getShortUrl", controllers.GetShortUrl)
}
