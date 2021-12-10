package main

import (
	"github.com/micro/message/email"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	addr := ":9000"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
	}

	server := email.Server{}

	grpcServer := grpc.NewServer()

	email.RegisterEmailServiceServer(grpcServer, &server)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
