package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "ok!")
	}
}
