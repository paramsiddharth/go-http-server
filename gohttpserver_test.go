package gohttpserver_test

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	gohttpserver "github.com/paramsiddharth/go-http-server"
)

func TestRun(t *testing.T) {
	var handlerObj interface{} = gohttpserver.ServerHandler{}
	if _, ok := handlerObj.(http.Handler); !ok {
		fmt.Println("Invalid handler!")
		os.Exit(1)
	}
}
