package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	exceptions "github.com/shortener-service/internal/common/exceptions"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) <= 0 {
			return
		}

		err := c.Errors.Last().Err

		if businessEx, ok := err.(*exceptions.BusinessException); ok {
			c.JSON(400, gin.H{
				"meta": gin.H{
					"code": businessEx.Code,
					"date": businessEx.DateTime,
				},
				"data": businessEx.Extra,
			})
			return
		}

		if databaseEx, ok := err.(*exceptions.DatabaseException); ok {
			c.JSON(501, gin.H{
				"meta": gin.H{
					"code": databaseEx.Code,
					"date": databaseEx.DateTime,
				},
				"data": gin.H{
					"message": "Service is unavailable for some time.",
				},
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": gin.H{
				"code":    "INTERNAL_SERVER_ERROR",
				"message": "internal server error",
			},
		})
	}
}
