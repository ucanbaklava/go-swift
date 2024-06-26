package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Get it from environment variable.
var jwtKey = []byte("secret_key")

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(username, role, email string) (string, error) {
	// TODO: get expiration time from environment variable or from a config file
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		Role:     role,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("error when generating jwt %w", err)
	}

	return signedToken, nil
}

func ParseJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(_ *jwt.Token) (any, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("error when parsing jwt %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("parseJWT => invalid token")
	}

	return claims, nil
}
