package repo

import (
	"database/sql"

	"github.com/jisu2000/go-auth-service/model"
	"github.com/jisu2000/go-auth-service/repo/query"
)

type UserRepo struct {
	DB *sql.DB
}

func (r *UserRepo) SaveUser(user *model.User) (model.User, error) {
	var createdUser model.User
	userRegisterQuery := query.USER_REGISTER_QUERY_AS_NORMAL_USER
	err := r.DB.QueryRow(
		userRegisterQuery,
		user.Name,
		user.Email,
		user.Password,
		user.MobileNumber).Scan(
		&createdUser.Id,
		&createdUser.CreatedAt,
		&createdUser.Name,
		&createdUser.Email,
		&createdUser.MobileNumber)
	return createdUser, err
}
