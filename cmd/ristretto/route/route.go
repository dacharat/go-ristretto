package route

import (
	"github.com/dacharat/go-ristretto/cmd/ristretto/handler"
	"github.com/dacharat/go-ristretto/pkg/caches"
	"github.com/gin-gonic/gin"
)

func NewRouter(c caches.ICache) *gin.Engine {
	route := gin.New()

	route.Use(gin.Recovery())

	h := handler.NewHandler(c)
	apiV1 := route.Group("/api/v1")
	{
		apiV1.GET("/:key", h.GetByKey)
		apiV1.POST("/:key", h.SetByKey)
	}

	return route
}
