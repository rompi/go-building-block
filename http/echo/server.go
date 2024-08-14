package echopackage

import (
	"time"

	httpserver "github.com/rompi/go-building-block/http"

	"github.com/labstack/echo/v4"
)

// EchoServer wraps the Echo instance.
type EchoServer struct {
	echo   *echo.Echo
	router *EchoRouter
}

// New creates and returns a new EchoServer instance.
func New() *EchoServer {
	e := echo.New()
	router := NewRouter(e)

	// Set up middleware
	e.Use(echo.MiddlewareFunc(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c)
			elapsed := time.Since(start)
			c.Logger().Infof("Request took %s", elapsed)
			return err
		}
	}))

	return &EchoServer{
		echo:   e,
		router: router,
	}
}

// AddRouter adds a router to the server.
func (s *EchoServer) AddRouter(router httpserver.Router) {
	s.router = router.(*EchoRouter)
}

// AddRouter adds a router to the server.
func (s *EchoServer) Router() httpserver.Router {
	return s.router
}

// Start starts the HTTP server on the given address.
func (s *EchoServer) Start(address string) error {
	return s.echo.Start(address)
}
