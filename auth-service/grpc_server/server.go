package main

import (
	"awesomeProject/config"
	"awesomeProject/grpc_server/auth"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func init() {
	config.InitializeEnvVars()
	config.Connect()
	config.ConnectRedis()
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50052")
	if err != nil {
		log.Fatal("Error running grpc_server!")
	}

	s := authorization.Server{}
	grpcServer := grpc.NewServer()

	authorization.RegisterAuthenticationServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Error running grpc grpc_server")
		panic("localhost:50052")
	} else {
		fmt.Println("localhost:50052")
		panic("localhost:50052")
	}
	panic("localhost:50052")

}
