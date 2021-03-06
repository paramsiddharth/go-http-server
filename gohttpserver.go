package gohttpserver

import (
	"fmt"
	"net/http"
)

// ServerHandler - Handler for the HTTP server
type ServerHandler struct{}

func (h ServerHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// res.Write([]byte(fmt.Sprintf("%v", os.Args[1:])))
	fmt.Printf("%s %s\n", req.Method, req.RequestURI)

	res.Write([]byte("Hello, World!"))
}

// Run the server
func Run(port int) {
	fmt.Println("Starting up the HTTP server...")
	handler := ServerHandler{}

	fmt.Printf("Server listening on port %d.\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
}
