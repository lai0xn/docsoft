package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lai0xn/docsoft/config"
	"github.com/lai0xn/docsoft/internal/models"
)

func GenerateToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"exp":    time.Now().Add(time.Hour * 100).Unix(),
	})
	tokenString, err := token.SignedString([]byte(config.JWTSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func RefreshToken(refreshToken string) (string, error) {
	// Parse and validate the refresh token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		// Check token signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid token signing method")
		}
		return config.JWTSecret, nil
	})
	if err != nil {
		return "", err
	}

	// Check if token is valid and not expired
	if !token.Valid {
		return "", errors.New("Invalid token")
	}

	// Extract user information from the token, such as user ID or username
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("Invalid token claims")
	}

	// Generate a new JWT token
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": claims["userID"],                        // Assuming user ID is stored in the refresh token claims
		"exp":    time.Now().Add(time.Minute * 15).Unix(), // Token expiration time
	})

	// Sign the token with the secret key
	tokenString, err := newToken.SignedString(config.JWTSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
