package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/brotherlogic/goserver"
	"google.golang.org/grpc"

	pb "github.com/brotherlogic/wink2rpc/proto"
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
func (r *HTTPRetriever) retrieve(url string, key string) []byte {
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		log.Fatalf("Error:%v", err)
	}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer: "+key)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	resp, err = client.Do(req)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

// DoRegister does RPC registration
func (s Server) DoRegister(server *grpc.Server) {
	pb.RegisterWinkServiceServer(server, &s)
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
