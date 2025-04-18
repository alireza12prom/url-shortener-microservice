package http

import (
	"github.com/gin-gonic/gin"
)

func SetupAccountRoutes(engine *gin.Engine, handler *AccountHandler) {
	router := engine.Group("/api/v1/account")

	{
		router.POST("/", handler.RegisterAccount)
		router.POST("/password.change", handler.ChangePassword)
		router.POST("/username.change", handler.ChangeUsername)

		router.GET("/", handler.GetAccount)
	}

	println("Account -> routes registered")
}
