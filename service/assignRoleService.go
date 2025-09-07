package service

import "github.com/jisu2000/go-auth-service/repo"

type AssignRoleService struct {
	REPO *repo.RoleRepo
}

func (srv *AssignRoleService) AttachUserRole(userId int) int {
	r, _ := srv.REPO.AssignUserRole(userId)
	return r
}
