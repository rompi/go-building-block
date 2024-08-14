package httpserver

type Router interface {
	GET(path string, handler interface{})
	POST(path string, handler interface{})
	PUT(path string, handler interface{})
	DELETE(path string, handler interface{})
	PATCH(path string, handler interface{})
	OPTIONS(path string, handler interface{})
	HEAD(path string, handler interface{})
}

// Server defines methods for starting an HTTP server.
type Server interface {
	Start(address string) error
	AddRouter(router Router)
	Router() Router
}
