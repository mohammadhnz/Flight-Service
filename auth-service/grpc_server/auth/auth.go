package authorization

import (
	"awesomeProject/controller"
	"awesomeProject/repository"
	"awesomeProject/utils"
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

	return &SignupResponse{
		Status: SignupResponse_OK_201, UserId: user.User_id,
	}, nil
}

func (s *Server) SignIn(ctx context.Context, signInData *SignInData) (*SignInResponse, error) {
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
		return &SignInResponse{
			Status:       SignInResponse_WRONG_EMAIL_OR_PHONE_NUMBER_OR_PASSWORD,
			AccessToken:  "",
			RefreshToken: "",
		}, nil
	}
	accessToken, refreshToken, err := utils.GetAuthTokens(user)

	if err != nil {
		return &SignInResponse{
			Status:       SignInResponse_WRONG_EMAIL_OR_PHONE_NUMBER_OR_PASSWORD,
			AccessToken:  "",
			RefreshToken: "",
		}, nil
	}

	return &SignInResponse{
		Status:       SignInResponse_OK_200,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Server) SignOut(ctx context.Context, signOutData *SignOutData) (*SignOutResponse, error) {
	log.Printf(
		"received signup request from client: access_token: %s refresh_token: %s",
		signOutData.AccessToken, signOutData.RefreshToken,
	)
	claims, user, err := utils.ExtractJwtToken(signOutData.AccessToken)
	if err != nil {
		return &SignOutResponse{Status: SignOutResponse_AUTHORIZATION_ERROR_401}, nil
	}
	err = controller.SignOutAndUpdateTokens(claims, user, signOutData.AccessToken)
	if err != nil {
		return &SignOutResponse{Status: SignOutResponse_FAILED_TO_SIGN_OUT}, nil
	}
	return &SignOutResponse{
		Status: SignOutResponse_OK_201,
	}, nil
}

func (s *Server) UserInfo(ctx context.Context, userInfoData *UserInfoData) (*UserInfoResponse, error) {
	log.Printf(
		"received signup request from client: access_token: %s refresh_token: %s",
		userInfoData.AccessToken, userInfoData.RefreshToken,
	)
	//claims, user, err := utils.ExtractJwtToken(signOutData.AccessToken)
	//if err != nil {
	//	return &UserInfoResponse{Status: UserInfoResponse_AUTHORIZATION_ERROR_401}, nil
	//}
	return &UserInfoResponse{Status: UserInfoResponse_OK_200}, nil
}
