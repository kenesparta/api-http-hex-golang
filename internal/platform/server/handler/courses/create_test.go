package courses

import (
	"api-http-hex-golang/internal/platform/storage/storagemocks"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
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
		createCourseRequest := createRequest{
			ID:   "833e025a-3e46-4249-945a-90de0609ba48",
			Name: "Demo Course",
		}

		b, err := json.Marshal(createCourseRequest)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/courses", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				require.NoError(t, err)
			}
		}(res.Body)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("given a valid request it returns 201", func(t *testing.T) {
		createCourseRequest := createRequest{
			ID:       "d25cc7bc-6758-4c58-9338-190d3f8b9742",
			Name:     "Demo Course",
			Duration: "10 months",
		}

		b, err := json.Marshal(createCourseRequest)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/courses", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				require.NoError(t, err)
			}
		}(res.Body)

		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}
