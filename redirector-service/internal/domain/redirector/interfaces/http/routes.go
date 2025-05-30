package http

import "github.com/gin-gonic/gin"

func SetupRedirectorRoutes(engine *gin.Engine, handler *RedirectorHandler) {
	base := engine.Group("/api/v1/redirector")

	{
		router := base.Group("/")

		router.GET("/:hash", handler.Redirect)
	}
}
