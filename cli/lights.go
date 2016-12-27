package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"google.golang.org/grpc"

	pbdi "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/wink2rpc/proto"
)

func getIP(servername string, ip string, port int) (string, int) {
	conn, _ := grpc.Dial(ip+":"+strconv.Itoa(port), grpc.WithInsecure())
	defer conn.Close()

	registry := pbdi.NewDiscoveryServiceClient(conn)
	entry := pbdi.RegistryEntry{Name: servername}
	r, _ := registry.Discover(context.Background(), &entry)
	return r.Ip, int(r.Port)
}

func main() {
	dServer, dPort := getIP("winkbridge", "10.0.1.17", 50055)
	dConn, err := grpc.Dial(dServer+":"+strconv.Itoa(dPort), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer dConn.Close()
	dClient := pb.NewWinkServiceClient(dConn)

	resp, err := dClient.ListDevices(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("Fatal error: %v", err)
	}

	for i, device := range resp.Device {
		fmt.Printf("%v. %v", i, device)
	}
}
