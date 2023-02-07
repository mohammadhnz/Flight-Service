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
	//
	//signupData := authorization.SignupData{
	//	PhoneNumber: "0933041sdf2310360",
	//	Email:       "dfsdfsdfsdfsdf@gmail.com",
	//	Password:    "Ali is near 999",
	//	Gender:      authorization.SignupData_WOMEN,
	//	FirstName:   "Mohammad123ali",
	//	LastName:    "Hossein123nezhad",
	//}
	//_, err = c.Signup(context.Background(), &signupData)
	//if err != nil {
	//	log.Printf("Could not sign up: %s", err)
	//}
	//
	//signInData := authorization.SignInData{
	//	PhoneNumber: "0933041sdf2310360",
	//	Email:       "dfsdfsdfsdfsdf@gmail.com",
	//	Password:    "Ali is near 999",
	//}

	signOutData := authorization.SignOutData{
		AccessToken:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzU3NzQ1MjMsInN1YiI6MTl9.c7MW30A1XK25IOGE3oFdBoSCHT7TgaSu7gW83eecWJw",
		RefreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzU4MTQxMjMsInN1YiI6MTl9.7kWka0V6yXmOqwz2cNctut7UEDNevy0l_Bbql5UEbXI",
	}
	res, err := c.SignOut(context.Background(), &signOutData)
	if err != nil {
		log.Printf("Could not sign up: %s", err)
	}
	log.Println(res)

	//res, err := c.SignIn(context.Background(), &signInData)
	//if err != nil {
	//	log.Printf("Could not sign up: %s", err)
	//}
	//log.Println(res)
}
