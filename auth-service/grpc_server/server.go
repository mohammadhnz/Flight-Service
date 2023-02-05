package main

import (
	"awesomeProject/grpc_server/auth"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal("Error running grpc_server!")
	}

	s := authorization.Server{}
	grpcServer := grpc.NewServer()

	authorization.RegisterAuthenticationServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Error running grpc grpc_server")
	}

}
