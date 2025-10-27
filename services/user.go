package services

import (
	"e-commerce/config"
	"e-commerce/models"
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(username, password string) (message string, err error) {
	var user models.User

	config.DB.Where("username =?", username).First(&user)

	log.Print("user: ", user)
	if user.Id != 0 {
		return "", errors.New("username is taken")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	createUser := models.User{
		Username: username,
		Password: string(hashedPassword),
	}

	config.DB.Create(&createUser)

	return "user created", nil
}

func Login(username, password string) (message string, err error) {
	var user models.User

	config.DB.Where("username =?", username).First(&user)

	if user.Id == 0 {
		return "", errors.New("username or password is incorrect")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("username or password is incorrect")
	}

	generatedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.Id,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	token, err := generatedToken.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return token, nil
}
