package main

import (
	"context"

	pb "github.com/brotherlogic/wink2rpc/proto"
)

// ListDevices calls out to list all the wink devices
func (s *Server) ListDevices(ctx context.Context, in *pb.Empty) (*pb.DeviceList, error) {
	return &pb.DeviceList{}, nil
}
