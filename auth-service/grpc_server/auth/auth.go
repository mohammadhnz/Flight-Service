package authorization

import (
	"awesomeProject/repository"
	"golang.org/x/net/context"
	"log"
)

type Server struct {
}

func (s *Server) Signup(ctx context.Context, signupData *SignupData) (*SignupResponse, error) {
	log.Printf(
		"received signup request from client: username: %s password: %s email: %s",
		signupData.FirstName, signupData.Email, signupData.Password,
	)

	buser := repository.UserData{
		Email:        signupData.Email,
		Phone_number: signupData.PhoneNumber,
		First_name:   signupData.FirstName,
		Last_name:    signupData.LastName,
		Password:     signupData.Password,
		Gender:       string(signupData.Gender),
	}
	user, err := repository.CreateUser(buser)
	if err != nil {
		return &SignupResponse{Status: SignupResponse_FAILED_TO_CREATE_USER_400}, err
	}
	return &SignupResponse{Status: SignupResponse_OK_201, UserId: user.User_id}, nil
}

func (s *Server) SignIn(ctx context.Context, signInData *SignInData) (*SignInResponse, error) {
	log.Printf(
		"received signup request from client: email: %s phone_number: %s email: %s",
		signInData.Email, signInData.PhoneNumber, signInData.Password,
	)

	return &SignInResponse{Status: SignInResponse_OK_200}, nil
}

func (s *Server) SignOut(ctx context.Context, signOutData *SignOutData) (*SignOutResponse, error) {
	log.Printf(
		"received signup request from client: access_token: %s refresh_token: %s",
		signOutData.AccessToken, signOutData.RefreshToken,
	)

	return &SignOutResponse{Status: SignOutResponse_OK_201}, nil
}

func (s *Server) UserInfo(ctx context.Context, userInfoData *UserInfoData) (*UserInfoResponse, error) {
	log.Printf(
		"received signup request from client: access_token: %s refresh_token: %s",
		userInfoData.AccessToken, userInfoData.RefreshToken,
	)

	return &UserInfoResponse{Status: UserInfoResponse_OK_200}, nil
}
