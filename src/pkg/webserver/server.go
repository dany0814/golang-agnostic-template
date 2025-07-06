package webserver

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"golang-agnostic-template/src/pkg/config"
	"golang-agnostic-template/src/pkg/logger"

	"golang-agnostic-template/src/application/actors/web"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine          *gin.Engine
	httpAddr        string
	shutdownTimeout time.Duration
}

func NewServer(ctx context.Context) (context.Context, Server) {
	srv := Server{
		engine:          gin.New(),
		httpAddr:        fmt.Sprintf("%s:%d", config.Params.WebHost, config.Params.WebPort),
		shutdownTimeout: config.Params.ShutdownTimeout,
	}
	return serverContext(ctx), srv
}

func (s *Server) Run(ctx context.Context) error {
	log.Println("Server running on", s.httpAddr)

	srv := &http.Server{
		Addr:    s.httpAddr,
		Handler: s.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server shut down", err)
		}
	}()

	<-ctx.Done()
	ctxShutDown, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
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

func (s *Server) Routes(ctx context.Context, log logger.ILogger) {
	groups := web.NewGroupRoutes(ctx, log)
	for _, group := range groups {
		groupRouter := s.engine.Group(group.Name)
		for _, r := range group.Paths {
			switch r.Method {
			case "POST":
				groupRouter.POST(r.Path, r.Handler)
			case "PUT":
				groupRouter.PUT(r.Path, r.Handler)
			case "PATCH":
				groupRouter.PATCH(r.Path, r.Handler)
			case "DELETE":
				groupRouter.DELETE(r.Path, r.Handler)
			case "GET":
				groupRouter.GET(r.Path, r.Handler)
			}
		}
	}
}
