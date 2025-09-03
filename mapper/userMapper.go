package mapper

import (
	"github.com/jisu2000/go-auth-service/dto/request"
	"github.com/jisu2000/go-auth-service/dto/response"
	"github.com/jisu2000/go-auth-service/model"
)

func ToUserModel(req request.UserRegisterRequest) model.User {
	return model.User{
		Name:         req.Name,
		Email:        req.Email,
		Password:     req.Password,
		MobileNumber: req.MobileNumber,
	}
}

func ToUserResponse(user model.User) response.UserResponse {
	return response.UserResponse{
		Name:         user.Name,
		Email:        user.Email,
		MobileNumber: user.MobileNumber,
		CreatedAt:    user.CreatedAt,
	}
}
