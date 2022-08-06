package authentication

import (
	"github.com/RistekCSUI/sistech-finpro/application"
	"github.com/RistekCSUI/sistech-finpro/shared"
	"github.com/RistekCSUI/sistech-finpro/shared/dto"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type (
	ViewService interface {
		BuildRegisterRequest(body *dto.RegisterDto, accessToken string) (*dto.RegisterRequest, error)
		RegisterUser(request dto.RegisterRequest) (*dto.RegisterResponse, error)
		Login(request dto.LoginRequest) (*dto.LoginResponse, error)
	}

	viewService struct {
		application application.Holder
		shared      shared.Holder
	}
)

func (v *viewService) BuildRegisterRequest(body *dto.RegisterDto, accessToken string) (*dto.RegisterRequest, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		return nil, err
	}

	body.Password = string(bytes)

	request := &dto.RegisterRequest{
		RegisterDto: *body,
		Token:       accessToken,
	}

	return request, nil
}

func (v *viewService) RegisterUser(request dto.RegisterRequest) (*dto.RegisterResponse, error) {
	id, err := v.application.AuthService.Insert(request)
	if err != nil {
		v.shared.Logger.WithFields(logrus.Fields{
			"access": request.Token,
		}).Errorf("failed to register user for username: %s", request.Username)
		return nil, err
	}

	response := &dto.RegisterResponse{
		ID:       id.(primitive.ObjectID),
		Username: request.Username,
	}

	return response, nil
}

func (v *viewService) Login(request dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := v.application.AuthService.FindUser(request.AccessToken, request.Username)
	if err != nil {
		v.shared.Logger.WithFields(logrus.Fields{
			"access": request.AccessToken,
		}).Errorf("failed to login user for username: %s, with error: %s", request.Username, err.Error())
		return nil, err
	}

	token, err := v.shared.JWT.GenerateToken(user.ID.String())

	response := &dto.LoginResponse{
		Username: user.Username,
		Token:    token,
	}

	return response, nil
}

func NewViewService(application application.Holder, shared shared.Holder) ViewService {
	return &viewService{
		application: application,
		shared:      shared,
	}
}
