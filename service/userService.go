package service

import (
	"database/sql"
	"errors"

	"github.com/jisu2000/go-auth-service/constants"
	"github.com/jisu2000/go-auth-service/dto/request"
	"github.com/jisu2000/go-auth-service/dto/response"
	"github.com/jisu2000/go-auth-service/mapper"
	"github.com/jisu2000/go-auth-service/model"
	"github.com/jisu2000/go-auth-service/repo"
	"github.com/jisu2000/go-auth-service/util"
)

type UserService struct {
	REPO      *repo.UserRepo
	ROLE_REPO *repo.RoleRepo
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
	if err != nil {
		return response.UserResponse{}, err
	}
	_, err = srv.ROLE_REPO.AssignUserRole(createdUser.Id)
	if err != nil {
		return response.UserResponse{}, err
	}
	resp := mapper.ToUserResponse(createdUser)
	resp.Roles = []string{constants.USER_ROLE_NAME}
	return resp, nil
}

func (srv *UserService) FetchAllUsers() ([]response.UserResponse, error) {
	var (
		users []model.User
		err   error
		roles []model.UserRole
	)
	users, err = srv.REPO.GetAllUsers()
	if err != nil {
		panic("Unable to Fetch User List")
	}

	respUser := make([]response.UserResponse, 0)

	for _, v := range users {
		uRes := mapper.ToUserResponse(v)
		roles = srv.ROLE_REPO.GetAllRoleOfUser(v.Id)
		rRes := mapRole(roles)
		uRes.Roles = rRes
		respUser = append(respUser, uRes)
	}

	return respUser, nil
}

func (srv *UserService) DeleteUser(id int) error {
	rows, err := srv.REPO.DeleteUserById(id)
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("user not found")
	}
	return nil
}

func mapRole(roles []model.UserRole) []string {
	res := make([]string, 0)
	for _, r := range roles {
		res = append(res, r.Role)
	}
	return res
}

func (srv *UserService) FetchUserFromEmailOrMobile(userName string) (*response.UserResponse, error) {
	user, err := srv.REPO.GetUserByEmailOrMobileNumber(userName)
	if err != nil {
		// return (nil, nil) when no rows found
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	resp := mapper.ToUserResponse(user)
	roles := srv.ROLE_REPO.GetAllRoleOfUser(user.Id)
	resp.Roles = mapRole(roles)
	return &resp, nil
}

func (srv *UserService) GetUserById(id int) (*response.UserResponse, error) {
	user, err := srv.REPO.GetUserById(id)
	if err != nil {
		return nil, err
	}
	resp := mapper.ToUserResponse(user)
	roles := srv.ROLE_REPO.GetAllRoleOfUser(user.Id)
	resp.Roles = mapRole(roles)
	return &resp, nil
}
