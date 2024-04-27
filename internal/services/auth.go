package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lai0xn/docsoft/config"
	"github.com/lai0xn/docsoft/internal/models"
	"github.com/lai0xn/docsoft/internal/types"
	"github.com/lai0xn/docsoft/storage/db"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct{}

func (s *Auth) GenerateToken(email string, password string) (string, error) {
	var user models.User
	db.DB.Where("email = ?", email).Find(&user)
	if user.Email == "" {
		return "", errors.New("Wrong Credentials")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("Wrong Credentials")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Email": user.Email,
		"Id":    user.ID,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(config.JWTSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *Auth) Signup(payload types.SignupPayload) error {
	enc_password, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 12)
	if err != nil {
		return err
	}
	user := models.User{
		First_Name:   payload.First_Name,
		Last_Name:    payload.Last_Name,
		Email:        payload.Email,
		Company_Name: payload.Company_Name,
		Password:     string(enc_password),
		Phone_Number: payload.Phone_Number,
	}
	err = db.DB.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}
