package services

import (
	"errors"
	internal "gin-demo/internal/models"

	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func InitAuthService(database *gorm.DB) *AuthService {

	database.AutoMigrate(&internal.User{})
	return &AuthService{
		db: database,
	}

}

func (a *AuthService) LoginService(email *string, password *string) (*internal.User, error) {
	if email == nil || password == nil {
		return nil, errors.New("Email/Password can't be empty")
	}

	var user internal.User
	if err := a.db.Where("email=?", email).Where("password=?").Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil

}

func (a *AuthService) RegisterService(email *string, password *string) (*internal.User, error) {
	if email == nil || password == nil {
		return nil, errors.New("Email/Password can't be empty")
	}

	var user internal.User
	user.Email = *email
	user.Password = *password
	if err := a.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil

}
