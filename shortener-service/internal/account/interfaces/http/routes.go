package http

import (
	"github.com/gin-gonic/gin"
	"github.com/shortener-service/internal/account/interfaces/http/middlewares"
)

func SetupAccountRoutes(engine *gin.Engine, handler *AccountHandler) {
	base := engine.Group("/api/v1/account")
	base.Use(middlewares.ErrorHandlerMiddleware())

	{
		router := base.Group("/")

		router.POST("/register", handler.RegisterAccount)
		router.POST("/login", handler.LoginAccount)
	}

	{
		router := base.Group("/")
		router.Use(middlewares.AuthMiddleware())

		router.GET("/info", handler.GetAccount)
		router.POST("/password.change", handler.ChangePassword)
		router.POST("/username.change", handler.ChangeUsername)
	}
}
