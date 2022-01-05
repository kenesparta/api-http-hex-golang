package courses

import (
	"api-http-hex-golang/internal/platform/storage/storagemocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestHandler_Create(t *testing.T) {
	cr := new(storagemocks.CourseRepository)
	cr.On(
		"Save",
		mock.Anything,
		mock.AnythingOfType("mooc.Course"),
	).Return(nil)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/course", CreateHandler(cr))

	t.Run("given an invalid request it returns 400", func(t *testing.T) {

	})

	t.Run("given a valid request it returns 201", func(t *testing.T) {

	})
}
