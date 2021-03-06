package main

import (
	"fmt"
	"os"
	"strconv"

	gohttpserver "github.com/paramsiddharth/go-http-server"
)

func main() {
	args := os.Args[1:]

	var port int
	if len(args) < 1 {
		port = 8080
	} else {
		parsedInt, err := strconv.Atoi(args[0])
		if err == nil {
			port = parsedInt
		} else {
			fmt.Fprintf(os.Stderr, "Invalid port number \"%s\"!\n", args[0])
			os.Exit(1)
		}
	}

	gohttpserver.Run(port)
}
