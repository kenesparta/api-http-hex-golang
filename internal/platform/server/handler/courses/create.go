package courses

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "ok!")
	}
}
