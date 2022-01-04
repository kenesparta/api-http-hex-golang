package courses

import (
	mooc "api-http-hex-golang/internal"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

func CreateHandler(cr mooc.CourseRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req createRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		course := mooc.NewCourse(req.ID, req.Name, req.Duration)

		if err := cr.Save(c, course); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.Status(http.StatusCreated)
	}
}
