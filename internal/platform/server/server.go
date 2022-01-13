package server

import (
	mooc "api-http-hex-golang/internal"
	"api-http-hex-golang/internal/platform/server/handler/health"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

const shutdownTime = time.Second * 5

type Server struct {
	engine   *gin.Engine
	httpAddr string
	cr       mooc.CourseRepository
}

func New(ctx context.Context, host string, port uint, cr mooc.CourseRepository) (context.Context, Server) {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),
		cr:       cr,
	}
	srv.registerRoutes()
	return serverContext(ctx), srv
}

func (s *Server) Run(ctx context.Context) error {
	log.Println("Running")
	// return s.engine.Run(s.httpAddr)
	srv := &http.Server{
		Addr:    s.httpAddr,
		Handler: s.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server shutdown ", err)
		}
	}()

	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(ctx, shutdownTime)
	
	defer cancel()

	return srv.Shutdown(ctxShutDown)
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c
		cancel()
	}()

	return ctx
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	// s.engine.POST("/course", courses.CreateHandler(s.cr))
}
