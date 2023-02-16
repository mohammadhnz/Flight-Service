package main

import (
	authorization "awesomeProject/grpc_server/auth"
	"context"
	"google.golang.org/grpc"
	"fmt"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Printf("Could not connect: %s", err)
	}

	defer conn.Close()

	c := authorization.NewAuthenticationClient(conn)

	signupData := authorization.SignupData{
		PhoneNumber: "0933041sdf2310360",
		Email:       "dfsdfsdfsdfsdf@gmail.com",
		Password:    "Ali is near 999",
		Gender:      authorization.SignupData_WOMEN,
		FirstName:   "Mohammad123ali",
		LastName:    "Hossein123nezhad",
	}
	_, err = c.Signup(context.Background(), &signupData)
	if err != nil {
		log.Printf("Could not sign up: %s", err)
	}
	fmt.Println("Fuck life")
}
