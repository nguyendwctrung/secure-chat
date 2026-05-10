package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/nguyendwctrung/secure-chat/internal/models"
	"github.com/nguyendwctrung/secure-chat/internal/repositories"
	"github.com/nguyendwctrung/secure-chat/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func Register(req models.RegisterRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user := models.User{
		ID: uuid.New(),
		Username: req.Username,
		Email: req.Email,
		PasswordHash: string(hashedPassword),
	}

	return repositories.CreateUser(&user)
}

func Login(req models.LoginRequest) (string, error) {
	user, err := repositories.GetUserByEmail(req.Email)

	if err != nil {
		return "", errors.New("Invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(req.Password),
	)

	if err != nil {
		return "", errors.New("Invalid email or password")
	}

	token, err := utils.GenerateToken(user.ID.String())

	if err != nil {
		return "", err
	}

	return token, nil
}