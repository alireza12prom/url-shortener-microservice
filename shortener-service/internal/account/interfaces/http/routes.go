package http

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(engine *gin.Engine) {
	handler := Handler{}

	router := engine.Group("/account")
	{
		router.POST("/", handler.RegisterAccount)
		router.POST("/password.change", handler.ChangePassword)
		router.POST("/username.change", handler.ChangeUsername)

		router.GET("/", handler.GetAccount)
	}

	println("Account -> routes registered")
}
