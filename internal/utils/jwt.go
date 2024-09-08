package utils

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type TokenType string

const (
	Access            TokenType = "access"
	Refresh           TokenType = "refresh"
	expriredTimeDay             = time.Hour * 24
	expriredTimeMonth           = time.Hour * 24 * 30
)

func createToken(tokenType TokenType, userID uint) (string, string, error) {
	var expriredTime = time.Now().Add(expriredTimeDay)
	if tokenType == Refresh {
		expriredTime = time.Now().Add(expriredTimeMonth)
	}
	jti := uuid.New().String()
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": userID,
		"exp": expriredTime.Unix(),
		"jit": jti,
		"typ": tokenType,
	})

	// Read private key
	signBytes, err := os.ReadFile(PrivateKeyPath)
	if err != nil {
		return "", "", err
	}

	// Parse private key
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return "", "", err
	}

	// Sign the token
	tokenString, err := token.SignedString(signKey)

	if err != nil {
		return "", "", err
	}
	return tokenString, jti, nil
}

func CreateTokenPair(userID uint) (string, string, string, error) {
	accessToken, _, err := createToken(Access, userID)
	if err != nil {
		return "", "", "", err
	}
	refreshToken, jti, err := createToken(Refresh, userID)
	if err != nil {
		return "", "", "", err
	}
	return accessToken, refreshToken, jti, nil
}

func ValidateToken(token string) (jwt.MapClaims, error) {
	verifiedKey, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is RSA
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			log.Panicln(ok)
			return nil, errors.New("unexpected signing method")
		}

		// Read public key
		verifyBytes, err := os.ReadFile(PublicKeyPath)
		if err != nil {
			return nil, err
		}

		// Parse public key
		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
		if err != nil {
			log.Panicln(err)
			return nil, errors.New("failed to parse public key")
		}

		return verifyKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !verifiedKey.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := verifiedKey.Claims.(jwt.MapClaims)
	if !ok {
		log.Panicln(ok)
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
