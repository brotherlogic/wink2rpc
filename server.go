package main

import (
	"encoding/json"
	"log"

	"golang.org/x/net/context"

	pb "github.com/brotherlogic/wink2rpc/proto"
)

// Retriever - Bridge for doing http requests
type Retriever interface {
	retrieve(url string) []byte
}

type jsonUnmarshaller interface {
	Unmarshal([]byte, interface{}) error
}
type prodUnmarshaller struct{}

func (jsonUnmarshaller prodUnmarshaller) Unmarshal(inp []byte, v interface{}) error {
	return json.Unmarshal(inp, v)
}

type device struct {
	Name string
}

type listDevicesResponse struct {
	Data []device
}

// ListDevices calls out to list all the wink devices
func (s *Server) ListDevices(ctx context.Context, in *pb.Empty) (*pb.DeviceList, error) {
	url := "https://api.wink.com/users/me/wink_devices"
	page := s.retr.retrieve(url)

	response := listDevicesResponse{}
	log.Printf("RESPONSE = %v", response)
	s.marshaller.Unmarshal(page, &response)

	list := &pb.DeviceList{}
	for _, device := range response.Data {
		list.Device = append(list.Device, &pb.Device{Name: device.Name})
	}

	return list, nil
}
