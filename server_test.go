package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	"golang.org/x/net/context"

	pb "github.com/brotherlogic/wink2rpc/proto"
)

type TestRetriever struct{}

func (retriever TestRetriever) retrieve(URL string) []byte {
	strippedURL := strings.Replace(strings.Replace(URL[21:], "?", "_", -1), "&", "_", -1)
	blah, err := os.Open("testdata/" + strippedURL)
	if err != nil {
		log.Printf("Error opening test file %v", err)
	}
	body, _ := ioutil.ReadAll(blah)
	return body
}

// Gets a test server that'll pull from local files rather than reading out
func getTestServer() Server {
	s := Server{}
	s.retr = TestRetriever{}
	s.key = "madeupkey"
	s.marshaller = prodUnmarshaller{}
	return s
}

func TestListDevices(t *testing.T) {
	s := getTestServer()
	list, err := s.ListDevices(context.Background(), &pb.Empty{})

	if err != nil {
		t.Errorf("Failure to list devices: %v", err)
	}

	//Test results should have one device called winner
	if len(list.Device) != 1 || list.Device[0].Name != "winner" {
		t.Errorf("Error in listing devices: %v", list)
	}
}
