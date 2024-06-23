package auth

import (
	"errors"
	"fmt"
	"go-fiber-postgres-template/pkg/user"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
)

func GenerateToken(tokenType string, jti string, u user.User) (string, error) {
	if tokenType != "access" && tokenType != "refresh" {
		return "", errors.New("wrong token type")
	}
	exp := time.Now().Add(time.Hour * 1).Unix()
	if tokenType == "refresh" {
		exp = time.Now().Add(time.Hour * 72).Unix()
	}
	if tokenType == "access" {
		jti = uuid.NewString()
	}
	claims := jwt.MapClaims{
		"user": u.ID,
		"type": tokenType,
		"exp":  exp,
		"jti":  jti,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(os.Getenv("APP_SECRET"))
	t, err := token.SignedString([]byte(os.Getenv("APP_SECRET")))
	if err != nil {
		return "", errors.New("error when create token")
	}

	return t, nil
}

func (r *RefreshToken) GenerateTokenPairs(u user.User) (string, string, error) {
	accessToken, _ := GenerateToken("access", "", u)
	refreshToken, err := GenerateToken("refresh", r.Jti, u)
	if err != nil {
		return "", "", errors.New("error when create token pairs")
	}
	return accessToken, refreshToken, nil
}
