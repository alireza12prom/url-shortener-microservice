package http

import (
	"github.com/gin-gonic/gin"
	http_middlewares "github.com/shortener-service/internal/common/middlewares/http"
)

func SetupShortenerRoutes(engine *gin.Engine, handler *ShortenerHandler) {
	base := engine.Group("/api/v1/shortener")
	base.Use(http_middlewares.ErrorHandlerMiddleware())

	{
		router := base.Group("/")
		router.Use(http_middlewares.AuthMiddleware())

		router.POST("/shorten.random", handler.ShortenRandom)
		router.POST("/shorten.custom", handler.ShortenCustom)
	}
}
