package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	coprocess "github.com/TykTechnologies/tyk-protobuf/bindings/go"
)

func main() {
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Listening...")
	s := grpc.NewServer()
	coprocess.RegisterDispatcherServer(s, &Dispatcher{})
	s.Serve(lis)
}
