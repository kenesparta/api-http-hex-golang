package server

import (
	mooc "api-http-hex-golang/internal"
	"api-http-hex-golang/internal/platform/server/handler/courses"
	"api-http-hex-golang/internal/platform/server/handler/health"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine   *gin.Engine
	httpAddr string
	cr       mooc.CourseRepository
}

func New(host string, port uint, cr mooc.CourseRepository) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),
		cr:       cr,
	}
	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println()
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/course", courses.CreateHandler(s.cr))
}
