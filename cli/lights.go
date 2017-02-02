package main

import (
	"context"
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

	res, err := dClient.ListDevices(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("Fatal error: %v", err)
	}
	for _, dev := range res.Device {
		log.Printf("%v - %v", dev.Name, dev.ObjectId)
	}

	_, err = dClient.Switch(context.Background(), &pb.LightChange{Dev: &pb.Device{Name: "Bedroom", ObjectId: "9785e88d-980f-4b48-8eb9-69dc6d9a7b43"}, State: true})
	if err != nil {
		log.Fatalf("Fatal error: %v", err)
	}
}
