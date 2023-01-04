package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/giovanars/desafio-programacao-fullstack/config"
)

type Endpoint interface {
	Handler(gin.IRouter, *config.Config, *Server) error
}

type Server struct {
	*http.Server
	Router *gin.Engine
	//TODO: Add db
}

func NewServer(address string) *Server {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(requestid.New())
	router.Use(cors.Default())

	return &Server{
		Router: router,
		Server: &http.Server{Addr: address},
	}
}

func Success(c *gin.Context, data interface{}) {
	/*if data == nil{
		c.Status(http.StatusOK)
		return
	}*/

	c.JSON(http.StatusOK, data)
}

func RouteNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"error": "route not found"})
}

func NotFound(c *gin.Context, msg string) {
	c.JSON(http.StatusNotFound, gin.H{"error": msg})
}

func BadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func (s *Server) Load(prefix string, config *config.Config, endpoints ...Endpoint) (gin.IRouter, error) {
	router := s.Router.Group(prefix)
	for _, endpoint := range endpoints {
		err := endpoint.Handler(router, config, s)
		if err != nil {
			return nil, err
		}
	}
	return router, nil
}

func (s *Server) Start() error {
	s.Router.NoRoute(RouteNotFound)
	s.Server.Handler = s.Router
	return s.Server.ListenAndServe()
}

func (s *Server) WaitSignal(ctx context.Context) (serverShutdown chan struct{}) {
	serverShutdown = make(chan struct{})

	go func(ctx context.Context, connsClosed chan struct{}) {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
		<-sig

		if err := s.Server.Shutdown(context.Background()); err != nil && err != http.ErrServerClosed {
			print(fmt.Errorf("HTTP server shutdown error", err, nil))
		}

		close(serverShutdown)
	}(ctx, serverShutdown)

	return
}
