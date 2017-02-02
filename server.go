package main

import (
	"encoding/json"
	"log"
	"strconv"

	"golang.org/x/net/context"

	pb "github.com/brotherlogic/wink2rpc/proto"
)

// Retriever - Bridge for doing http requests
type Retriever interface {
	retrieve(url string, key string) []byte
	put(url string, key string, data string)
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
	UUID string
}

type listDevicesResponse struct {
	Data []device
}

// ListDevices calls out to list all the wink devices
func (s *Server) ListDevices(ctx context.Context, in *pb.Empty) (*pb.DeviceList, error) {
	url := "https://api.wink.com/users/me/wink_devices"
	page := s.retr.retrieve(url, s.key)

	response := listDevicesResponse{}
	log.Printf("RESPONSE = %v", page)
	s.marshaller.Unmarshal(page, &response)

	list := &pb.DeviceList{}
	for _, device := range response.Data {
		list.Device = append(list.Device, &pb.Device{Name: device.Name, ObjectId: device.UUID})
	}

	return list, nil
}

// Switch toggles a device
func (s *Server) Switch(ctx context.Context, state *pb.LightChange) (*pb.Empty, error) {
	url := "https://api.wink.com/light_bulbs/" + state.Dev.ObjectId
	s.retr.put(url, s.key, "{\"desired_state\": {\"powered\": "+strconv.FormatBool(state.State)+"}}")
	return &pb.Empty{}, nil
}
