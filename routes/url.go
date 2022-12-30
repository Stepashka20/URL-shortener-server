package routes

import (
	"stepashka20/url-shortener/controllers"

	"github.com/gin-gonic/gin"
)

type routerEngine struct {
	*gin.Engine
}

func Init() *routerEngine {
	return &routerEngine{gin.Default()}
}

func (router *routerEngine) UrlRoute() {
	router.GET("/:key", controllers.GetRedirect)
	router.POST("/getShortUrl", controllers.GetShortUrl)
}
