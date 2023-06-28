package main

import (
	"context"
	"fmt"
	"github.com/Odery/gRPC-Playground/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
	"net"
)

type CustomerServer struct {
	pb.UnimplementedCustomerCreatorServer
}

func (_ CustomerServer) CreateCustomer(_ context.Context, nc *pb.NewCustomer) (*pb.Customer, error) {
	log.Printf("[INFO]: New Customer recieved: %v\n", nc)
	if nc.Name != "" && nc.LastName != "" {
		comment := "test Customer"
		customer := &pb.Customer{
			ID:        uuid.New().ID(),
			Name:      nc.Name,
			LastName:  nc.LastName,
			Age:       nc.Age,
			AutoLogin: "pseudo login link",
			Comment:   &comment,
		}
		return customer, nil
	} else {
		return nil, fmt.Errorf("incorect Customer information")
	}
}

func main() {
	s, err := net.Listen("tcp", "0.0.0.0:8089")
	if err != nil {
		log.Fatalf("failed to start TCP listenere: %s", err)
	}

	server := grpc.NewServer()
	service := new(CustomerServer)

	pb.RegisterCustomerCreatorServer(server, service)

	err = server.Serve(s)
	if err != nil {
		log.Fatal("failed to server Customer service")
	}
}
