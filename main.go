package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/rompi/go-building-block/logger"

	echopackage "github.com/rompi/go-building-block/http/echo"

	"github.com/labstack/echo/v4"
)

// ExampleHandler is a sample handler function.
func ExampleHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Example Endpoint")
}

// ExampleHandler is a sample handler function.
func ExampleHandler2(c echo.Context) error {
	return c.String(http.StatusOK, "Example Endpoint2")
}

func main() {
	// call logger.NewLogger()
	var buf bytes.Buffer
	log := logger.NewLogger(&buf, nil)

	// call logger.Info()
	log.Info("This is an info message")
	// call logger.Warning()
	log.Warning("This is a warning message")
	// call logger.Error()
	log.Error("This is an error message")

	server := echopackage.New()
	// Define routes
	server.AddRouter(server.Router()) // Pass the router to the server
	server.Router().GET("/example", ExampleHandler)
	server.Router().POST("/example", ExampleHandler2)

	// go func() {
	if err := server.Start(":8080"); err != nil {
		log.Error(fmt.Sprintf("Error starting server: %v", err))
	}
	// }()
}
