package repo

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/jisu2000/go-auth-service/model"
	"github.com/jisu2000/go-auth-service/repo/query"
)

type RefreshTokenRepo struct {
	DB *sql.DB
}

func (r *RefreshTokenRepo) SaveRefreshToken(userId int) (model.RefreshToken, error) {
	newRefreshToken := uuid.New()
	token := newRefreshToken.String()
	var insertedToken string

	// Using QueryRow with RETURNING token to capture inserted token
	err := r.DB.QueryRow(query.INSERT_RERESHTOKEN_QUERY, token, userId).Scan(&insertedToken)
	return model.RefreshToken{
		Token:  insertedToken,
		UserId: userId,
	}, err
}

func (r *RefreshTokenRepo) FetchRefreshToken(token string) (*model.RefreshToken, error) {
	var refreshToken model.RefreshToken
	row := r.DB.QueryRow(query.FETCH_REFRESH_TOKEN_BY_TOKEN_QUERY, token)
	if err := row.Scan(&refreshToken.Id, &refreshToken.Token, &refreshToken.UserId, &refreshToken.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &refreshToken, nil
}
