package authorization

import (
	"awesomeProject/config"
	"awesomeProject/controller"
	"awesomeProject/repository"
	"awesomeProject/utils"
	"fmt"
	"golang.org/x/net/context"
	"log"
	"os"
	"time"
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

	return &SignupResponse{
		Status: SignupResponse_OK_201, UserId: user.User_id,
	}, nil
}

func (s *Server) SignIn(ctx context.Context, signInData *SignInData) (*FreshTokensResponse, error) {
	log.Printf(
		"received signup request from client: email: %s phone_number: %s email: %s",
		signInData.Email, signInData.PhoneNumber, signInData.Password,
	)
	buser := repository.GetUserData{
		Email:        signInData.Email,
		Phone_number: signInData.PhoneNumber,
		Password:     signInData.Password,
	}

	user, err := repository.GetUser(buser)

	if err != nil {
		return &FreshTokensResponse{
			Status:       FreshTokensResponse_WRONG_EMAIL_OR_PHONE_NUMBER_OR_PASSWORD,
			AccessToken:  "",
			RefreshToken: "",
		}, nil
	}
	accessToken, refreshToken, err := utils.GetAuthTokens(user)

	if err != nil {
		return &FreshTokensResponse{
			Status:       FreshTokensResponse_WRONG_EMAIL_OR_PHONE_NUMBER_OR_PASSWORD,
			AccessToken:  "",
			RefreshToken: "",
		}, nil
	}

	return &FreshTokensResponse{
		Status:       FreshTokensResponse_OK_200,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Server) SignOut(ctx context.Context, tokenData *TokenData) (*SignOutResponse, error) {
	log.Printf(
		"received signup request from client: access_token: %s refresh_token: %s",
		tokenData.AccessToken, tokenData.RefreshToken,
	)
	err, response, err2, done := checkIfTokenHasBeenExpired(tokenData)
	if done {
		return response, err2
	}
	claims, user, err := utils.ExtractJwtToken(tokenData.AccessToken, os.Getenv("ACCESS_SECRET"))
	if err != nil {
		return &SignOutResponse{Status: SignOutResponse_AUTHORIZATION_ERROR_401}, nil
	}
	err = controller.SignOutAndUpdateTokens(claims, user, tokenData.AccessToken)
	if err != nil {
		return &SignOutResponse{Status: SignOutResponse_FAILED_TO_SIGN_OUT}, nil
	}
	config.RedisClient.Set(tokenData.AccessToken, "Some thing", 24*time.Hour)
	return &SignOutResponse{
		Status: SignOutResponse_OK_201,
	}, nil
}

func checkIfTokenHasBeenExpired(tokenData *TokenData) (error, *SignOutResponse, error, bool) {
	data, err := config.RedisClient.Get(tokenData.AccessToken).Result()
	if data != "" {
		return nil, &SignOutResponse{
			Status: SignOutResponse_OK_201,
		}, nil, true

	}
	if err != nil {
		_, _, err := utils.CheckTokenIsNotUnAuthorized(tokenData.AccessToken)
		if err != nil {
			return err, nil, nil, false
		}

	}
	return err, nil, nil, false
}

func (s *Server) UserInfo(ctx context.Context, tokenData *TokenData) (*UserInfoResponse, error) {
	log.Printf(
		"received signup request from client: access_token: %s refresh_token: %s",
		tokenData.AccessToken, tokenData.RefreshToken,
	)
	_, user, err := utils.ExtractJwtToken(tokenData.AccessToken, os.Getenv("ACCESS_SECRET"))
	if err != nil {
		return &UserInfoResponse{Status: UserInfoResponse_AUTHORIZATION_ERROR_401}, nil
	}
	return &UserInfoResponse{
		Status: UserInfoResponse_OK_200,
		User: &UserData{
			Id:          user.User_id,
			FirstName:   user.First_name,
			LastName:    user.Last_name,
			PhoneNumber: user.Phone_number,
			Email:       user.Email,
			Gender:      user.Gender,
		},
	}, nil
}
func (s *Server) Refresh(ctx context.Context, tokenData *TokenData) (*FreshTokensResponse, error) {
	log.Printf(
		"received signup request from client: access_token: %s refresh_token: %s",
		tokenData.AccessToken, tokenData.RefreshToken,
	)
	_, user, err := utils.ExtractJwtToken(tokenData.RefreshToken, os.Getenv("REFRESH_SECRET"))
	if err != nil {
		fmt.Println("Err 1")
		return &FreshTokensResponse{
			Status:       FreshTokensResponse_FAILED_TO_CREATE_NEW_TOKEN,
			AccessToken:  tokenData.AccessToken,
			RefreshToken: tokenData.RefreshToken,
		}, nil
	}
	accessToken, refreshToken, err := utils.GetAuthTokens(user)
	if err != nil {
		fmt.Println("Err 2")
		return &FreshTokensResponse{
			Status:       FreshTokensResponse_FAILED_TO_CREATE_NEW_TOKEN,
			AccessToken:  tokenData.AccessToken,
			RefreshToken: tokenData.RefreshToken,
		}, nil
	}
	fmt.Println("Err 3")
	return &FreshTokensResponse{
		Status:       FreshTokensResponse_OK_200,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
