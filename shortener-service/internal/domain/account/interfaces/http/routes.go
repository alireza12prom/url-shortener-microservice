package http

import (
	"github.com/gin-gonic/gin"
	http_middlewares "github.com/shortener-service/internal/common/middlewares/http"
)

func SetupAccountRoutes(engine *gin.Engine, handler *AccountHandler) {
	base := engine.Group("/api/v1/account")
	base.Use(http_middlewares.ErrorHandlerMiddleware())

	{
		router := base.Group("/")

		router.POST("/register", handler.RegisterAccount)
		router.POST("/login", handler.LoginAccount)
	}

	{
		router := base.Group("/")
		router.Use(http_middlewares.AuthMiddleware())

		router.GET("/info", handler.GetAccount)
		router.POST("/password.change", handler.ChangePassword)
		router.POST("/username.change", handler.ChangeUsername)
	}
}
