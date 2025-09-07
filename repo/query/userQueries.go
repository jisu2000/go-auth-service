package query

const (
	USER_REGISTER_QUERY_AS_NORMAL_USER string = `INSERT INTO users 
				(name, email, password, mobile_number) 
				VALUES ($1, $2, $3, $4) 
				RETURNING id, created_at, name, email, mobile_number`

	FETCH_ALL_USER_QUERY string = `SELECT id,created_at,name,email,mobile_number FROM users`

	ADD_ROLE_QUERY string = `INSERT INTO user_roles
				(user_id,role_id)
				VALUES($1,$2)`

	FETCH_ALL_ROLE_QUERY string = `SELECT r.id,r.role FROM roles r
				INNER JOIN user_roles ur
				ON r.id = ur.role_id
				WHERE ur.user_id = $1;`

	DELETE_USER_QUERY string = `DELETE FROM users WHERE id = $1`

	FETCH_USER_BY_EMAIL_OR_MOBILE_NUMBER_QUERY string = `SELECT id,created_at,name,email,password,mobile_number FROM users 
                WHERE email = $1 OR mobile_number = $1`

	FETCH_USER_BY_ID_QUERY string = `SELECT id,created_at,name,email,mobile_number FROM users WHERE id = $1`
)
