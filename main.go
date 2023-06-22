package main

import (
	"github.com/Odery/gRPC-Playground/pb"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
)

func main() {
	p := &pb.Person{
		ID:       10,
		Name:     "Ivan",
		LastName: "Reznikov",
		Age:      33,
		Email:    "reznikov@mial.net",
		Comment:  nil,
	}

	data, err := proto.Marshal(p)
	if err != nil {
		log.Println(err)
	}

	err = os.WriteFile("proto.data", data, 0644)
	if err != nil {
		log.Println(err)
	}
}
