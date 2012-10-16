package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/dqminh/gocached"
	"io"
	"net"
)

var port *int = flag.Int("port", 11212, "Port on which to listen")

type Server struct {
	server       *gocached.Cache
	requestQueue chan *Context
}

func (server *Server) Start() {
	for {
		ctx := <-server.requestQueue
		log.Println("Got a context: %s", ctx)
		server.handleRequest(ctx)
		server.requestQueue <- ctx
	}
}

// listen and process request
func listenToRequest(ln) {
	requestQueue := make(chan RequestContext)
	server := &Server{server: gocached.NewCached(), requestQueue: requestQueue}

	go server.Start()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting request.", err)
		} else {
			fmt.Println("Accept connection from address: %v", err.RemoteAddr())
			// do something to handle the connection
		}
	}
}

// Main function of the server
// Responsible for starting the server and accepts request
func main() {
	fmt.Println("Starting server at port %d", port)
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("Error starting server at port: 12345 with error: %s", err)
	}
	listenToRequest(ln)

}
