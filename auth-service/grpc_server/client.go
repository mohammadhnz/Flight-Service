package main

import (
	authorization "awesomeProject/grpc_server/auth"
	"context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Printf("Could not connect: %s", err)
	}

	defer conn.Close()

	c := authorization.NewAuthenticationClient(conn)

	signupData := authorization.SignupData{
		PhoneNumber: "09330410360",
		Email:       "hmohammadali2013@gmail.com",
		Password:    "Ali is near 999",
		Gender:      authorization.SignupData_WOMEN,
		FirstName:   "Mohammadali",
		LastName:    "Hosseinnezhad",
	}
	response, err := c.Signup(context.Background(), &signupData)
	if err != nil {
		log.Printf("Could not sign up: %s", err)
	}

	log.Printf("Response status from grpc_server: %s", response.Status)
}
