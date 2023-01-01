package routes

import (
	"stepashka20/url-shortener/controllers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
)

type routerEngine struct {
	*gin.Engine
}

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.String(429, "Too many requests.")
}

var mw gin.HandlerFunc

func Init() *routerEngine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "GET", "OPTIONS"},
	}))
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Second,
		Limit: 3,
	})
	mw = ratelimit.RateLimiter(store, &ratelimit.Options{
		KeyFunc:      keyFunc,
		ErrorHandler: errorHandler,
	})
	// r.Use(mw)
	return &routerEngine{r}
}

func (router *routerEngine) UrlRoute() {
	router.GET("/:key", controllers.GetRedirect)
	router.POST("/getShortUrl", mw, controllers.GetShortUrl)
}
