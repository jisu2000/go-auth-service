package service

import (
	"github.com/jisu2000/go-auth-service/dto/request"
	"github.com/jisu2000/go-auth-service/dto/response"
	"github.com/jisu2000/go-auth-service/mapper"
	"github.com/jisu2000/go-auth-service/model"
	"github.com/jisu2000/go-auth-service/repo"
	"github.com/jisu2000/go-auth-service/util"
)

type UserService struct {
	REPO *repo.UserRepo
}

func (srv *UserService) CreateNewUser(request request.UserRegisterRequest) (response.UserResponse, error) {
	var modelUser model.User
	passwordHash, err := util.HashPassword(request.Password)
	if err != nil {
		panic("Error hashing password")
	}
	modelUser = mapper.ToUserModel(request)
	modelUser.Password = passwordHash
	createdUser, err := srv.REPO.SaveUser(&modelUser)
	return mapper.ToUserResponse(createdUser), err
}
