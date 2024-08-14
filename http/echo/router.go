package echopackage

import "github.com/labstack/echo/v4"

// EchoRouter wraps the Echo instance and implements the Router interface.
type EchoRouter struct {
	echo *echo.Echo
}

// NewRouter creates a new EchoRouter instance.
func NewRouter(e *echo.Echo) *EchoRouter {
	return &EchoRouter{echo: e}
}

// GET adds a GET route to the Echo router.
func (r *EchoRouter) GET(path string, handler interface{}) {
	r.echo.GET(path, handler.(func(c echo.Context) error))
}

// POST adds a POST route to the Echo router.
func (r *EchoRouter) POST(path string, handler interface{}) {
	r.echo.POST(path, handler.(func(c echo.Context) error))
}

// PUT adds a PUT route to the Echo router.
func (r *EchoRouter) PUT(path string, handler interface{}) {
	r.echo.PUT(path, handler.(func(c echo.Context) error))
}

// DELETE adds a DELETE route to the Echo router.
func (r *EchoRouter) DELETE(path string, handler interface{}) {
	r.echo.DELETE(path, handler.(func(c echo.Context) error))
}

// PATCH adds a PATCH route to the Echo router.
func (r *EchoRouter) PATCH(path string, handler interface{}) {
	r.echo.PATCH(path, handler.(func(c echo.Context) error))
}

// OPTIONS adds an OPTIONS route to the Echo router.
func (r *EchoRouter) OPTIONS(path string, handler interface{}) {
	r.echo.OPTIONS(path, handler.(func(c echo.Context) error))
}

// HEAD adds a HEAD route to the Echo router.
func (r *EchoRouter) HEAD(path string, handler interface{}) {
	r.echo.HEAD(path, handler.(func(c echo.Context) error))
}
