package repo

import (
	"database/sql"

	"github.com/jisu2000/go-auth-service/constants"
	"github.com/jisu2000/go-auth-service/model"
	"github.com/jisu2000/go-auth-service/repo/query"
)

type RoleRepo struct {
	DB *sql.DB
}

func (r *RoleRepo) AssignUserRole(userId int) (int, error) {
	assignRoleQuery := query.ADD_ROLE_QUERY
	res, err := r.DB.Exec(
		assignRoleQuery,
		userId,
		constants.USER_ROLE_ID,
	)
	if err != nil {
		return 0, err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		return 0, err
	}

	return int(rows), nil

}

func (r *RoleRepo) GetAllRoleOfUser(userId int) []model.UserRole {
	roles := make([]model.UserRole, 0)

	rows, err := r.DB.Query(query.FETCH_ALL_ROLE_QUERY, userId)

	if err != nil {
		return roles
	}

	for rows.Next() {
		var role model.UserRole

		if err := rows.Scan(&role.Id, &role.Role); err != nil {
			return roles
		}
		roles = append(roles, role)
	}

	return roles

}
