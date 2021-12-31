package courses

import (
	mooc "api-http-hex-golang/internal"
	"api-http-hex-golang/internal/platform/storage/mysql"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	dbUser = ""
	dbPass = ""
	dbHost = ""
	dbPort = ""
	dbName = ""
)

type createRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

func CreateHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req createRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		course := mooc.NewCourse(req.ID, req.Name, req.Duration)

		mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
		db, err := sql.Open("mysql", mysqlURI)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		courseRepository := mysql.NewCourseRepository(db)

		if err := courseRepository.Save(c, course); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.Status(http.StatusCreated)
	}
}
