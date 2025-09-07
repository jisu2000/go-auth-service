package query

const (
	INSERT_RERESHTOKEN_QUERY string = `INSERT INTO refresh_tokens (token, user_id) VALUES ($1, $2) RETURNING token`

	FETCH_REFRESH_TOKEN_BY_TOKEN_QUERY string = `SELECT id, token, user_id,created_at FROM refresh_tokens WHERE token = $1`

	DELETE_REFRESH_TOKEN_BY_USER_ID_QUERY string = `DELETE FROM refresh_tokens WHERE user_id = $1`
)
