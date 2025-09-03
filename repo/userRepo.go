package repo

import (
	"database/sql"
	"log"

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

func (r *UserRepo) GetAllUsers() ([]model.User, error) {
	rows, err := r.DB.Query(query.FETCH_ALL_USER_QUERY)

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	defer rows.Close()

	users := make([]model.User, 0)

	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.Id, &u.CreatedAt, &u.Name, &u.Email, &u.MobileNumber); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil

}
