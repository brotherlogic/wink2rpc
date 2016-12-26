package main

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/wink2rpc/proto"
)

func TestListDevices(t *testing.T) {
	s := Server{}
	_, err := s.ListDevices(context.Background(), &pb.Empty{})

	if err != nil {
		t.Errorf("Failure to list devices: %v", err)
	}
}
