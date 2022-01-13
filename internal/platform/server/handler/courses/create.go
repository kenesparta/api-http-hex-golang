package courses

import (
	mooc "api-http-hex-golang/internal"
	"api-http-hex-golang/internal/creating"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

// CreateHandler Representar el caso de uso de fora atomica, logia agnostica, encapsular hacia el Application Service.
func CreateHandler(cs creating.CourseService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req createRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err := cs.CreateCourse(c, req.ID, req.Name, req.Duration)

		if err != nil {
			switch {
			case errors.Is(err, mooc.ErrInvalidCourseID):
				c.JSON(http.StatusBadRequest, err.Error())
				return
			default:
				c.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		c.Status(http.StatusCreated)
	}
}
