package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"google.golang.org/grpc"

	"github.com/brotherlogic/goserver"
)

// Server the configuration for the syncer
type Server struct {
	*goserver.GoServer
	key        string
	retr       Retriever
	marshaller jsonUnmarshaller
}

// HTTPRetriever pulls http pages
type HTTPRetriever struct{}

// Does a web retrieve
func (r *HTTPRetriever) retrieve(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

// DoRegister does RPC registration
func (s Server) DoRegister(server *grpc.Server) {
	// Does nothing
}

// InitServer builds an initial server
func InitServer(key *string) Server {
	server := Server{&goserver.GoServer{}, *key, &HTTPRetriever{}, &prodUnmarshaller{}}
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
