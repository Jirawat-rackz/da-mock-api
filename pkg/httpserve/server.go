package httpserve

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// IServer is an interface that defines the methods that a server must implement
type IServer interface {
	// Start starts the server
	Start() error

	// Use adds middleware to the server
	Use(middleware ...gin.HandlerFunc) gin.IRoutes

	// Group creates a new router group
	Group(relativePath string, handlers ...gin.HandlerFunc) *gin.RouterGroup

	// GET adds a GET route to the server
	GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes

	// POST adds a POST route to the server
	POST(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes

	// PUT adds a PUT route to the server
	PUT(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes

	// DELETE adds a DELETE route to the server
	DELETE(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes

	// PATCH adds a PATCH route to the server
	PATCH(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
}

// Server is a struct that implements the IServer interface
type Server struct {
	Engine *gin.Engine
	Port   string
}

func New(port string) IServer {
	return &Server{
		Engine: gin.Default(),
		Port:   port,
	}
}

func (s Server) Start() error {

	var (
		err error
	)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", s.Port),
		Handler: s.Engine,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Server is running on port %s", s.Port)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")

	return err
}
