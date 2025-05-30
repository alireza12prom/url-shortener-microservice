package http

import (
	"github.com/gin-gonic/gin"
	http_middlewares "github.com/shortener-service/internal/common/middlewares/http"
)

func SetupRedirectorRoutes(engine *gin.Engine, handler *RedirectorHandler) {
	base := engine.Group("/api/v1/redirector")
	base.Use(http_middlewares.ErrorHandlerMiddleware())

	{
		router := base.Group("/")

		router.GET("/redirect/:hash", handler.Redirect)
	}
}
