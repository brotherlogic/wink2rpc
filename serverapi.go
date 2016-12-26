package main

import (
	"flag"

	"google.golang.org/grpc"

	"github.com/brotherlogic/goserver"
)

// Server the configuration for the syncer
type Server struct {
	*goserver.GoServer
	key string
}

// DoRegister does RPC registration
func (s Server) DoRegister(server *grpc.Server) {
	// Does nothing
}

// InitServer builds an initial server
func InitServer(key *string) Server {
	server := Server{&goserver.GoServer{}, *key}
	server.Register = server

	return server
}

func main() {
	var key = flag.String("key", "", "OAuth key for wink API server.")
	flag.Parse()

	server := InitServer(key)

	server.PrepServer()
	server.RegisterServer("winkbridge", false)
	server.Serve()
}
