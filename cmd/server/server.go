package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/mauriciolucas22/gRPC-examples/pb"
	"github.com/mauriciolucas22/gRPC-examples/services"
)

func main() {

	lis, err := net.Listen("tcp", "localhost:5000")

	if err != nil {
		log.Fatalf("Error on server connection: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not server: %v", err)
	}
}
