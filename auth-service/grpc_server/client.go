package main

import (
	authorization "awesomeProject/grpc_server/auth"
	"context"
	"google.golang.org/grpc"
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

	signInData := authorization.SignInData{
		PhoneNumber: "0933041sdf2310360",
		Email:       "dfsdfsdfsdfsdf@gmail.com",
		Password:    "Ali is near 999",
	}

	res, err := c.SignIn(context.Background(), &signInData)
	if err != nil {
		log.Printf("Could not sign up: %s", err)
	}
	log.Println(res)
	userInfoData := authorization.UserInfoData{
		AccessToken:  string(res.AccessToken),
		RefreshToken: string(res.RefreshToken),
	}
	res1, err1 := c.UserInfo(context.Background(), &userInfoData)
	if err1 != nil {
		log.Printf("Could not sign up: %s", err)
	}
	log.Println(res1)
}
