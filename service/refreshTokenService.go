package service

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jisu2000/go-auth-service/repo"
	"github.com/joho/godotenv"
)

type RefreshTokenService struct {
	USERREPO    *repo.UserRepo
	REFRESHREPO *repo.RefreshTokenRepo
}

func (srv *RefreshTokenService) CreateNewRefreshToken(userId int) (*string, error) {
	fetchedUser, err := srv.USERREPO.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	refreshToken, err := srv.REFRESHREPO.SaveRefreshToken(fetchedUser.Id)
	return &refreshToken.Token, err
}

func (srv *RefreshTokenService) VerifyRefreshToken(refreshToken string) (*string, *int, error) {
	fetchedToken, err := srv.REFRESHREPO.FetchRefreshToken(refreshToken)

	if err != nil || fetchedToken == nil {
		return nil, nil, err
	}

	fetchedUser, err := srv.USERREPO.GetUserById(fetchedToken.UserId)
	if err != nil {
		return nil, nil, err
	}
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}
	createdAtTime, err := time.Parse(time.RFC3339, fetchedToken.CreatedAt)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid CreatedAt format: %w", err)
	}
	ttlDays, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_TTL_DAYS"))
	expirationTime := createdAtTime.Add(time.Hour * 24 * time.Duration(ttlDays))

	if time.Now().After(expirationTime) {
		return nil, nil, fmt.Errorf("refresh token has expired")
	}
	return &fetchedToken.Token, &fetchedUser.Id, nil
}
