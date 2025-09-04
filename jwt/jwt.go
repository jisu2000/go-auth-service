package jwt

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Sub   string
	Roles []string
	jwt.RegisteredClaims
}

func CreateToken(email string, roles []string) (string, time.Time, error) {

	secret := []byte(os.Getenv("JWT_SECRET"))
	ttlMin, _ := strconv.Atoi(os.ExpandEnv("JWT_TTL_MINUTES"))

	if ttlMin <= 0 {
		ttlMin = 60
	}

	exp := time.Now().Add(time.Duration(ttlMin) * time.Minute)

	claims := Claims{
		Sub:   email,
		Roles: roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := t.SignedString(secret)
	return s, exp, err

}

func ParseToken(tokenString string) (*Claims, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil

}
