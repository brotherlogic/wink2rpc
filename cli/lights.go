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

	_, err = dClient.Switch(context.Background(), &pb.LightChange{Dev: &pb.Device{Name: "Bedroom"}, State: true})
	if err != nil {
		log.Fatalf("Fatal error: %v", err)
	}
}
