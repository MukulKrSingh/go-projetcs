package services

import (
	"errors"
	"fmt"
	internal "gin-demo/internal/models"
	"gin-demo/internal/utils"

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

func (a *AuthService) CheckUserExistOrNot(email *string) bool {
	var user internal.User
	if err := a.db.Where("email=?", email).Find(&user).Error; err != nil {
		return false
	}
	if user.Email != "" {
		return true
	}
	return false

}

func (a *AuthService) LoginService(email *string, password *string) (*internal.User, error) {
	if email == nil || password == nil {
		return nil, errors.New("Email/Password can't be empty")
	}

	var user internal.User
	if err := a.db.Where("email=?", email).Find(&user).Error; err != nil {
		return nil, err
	}
	if user.Email == "" {
		return nil, errors.New("no user found")

	}

	if !utils.CheckPasswordHash(*password, user.Password) {
		return nil, errors.New("incorrect password")
	}
	fmt.Println(utils.CheckPasswordHash(*password, user.Password))
	return &user, nil

}

func (a *AuthService) RegisterService(email *string, password *string) (*internal.User, error) {
	if email == nil || password == nil {
		return nil, errors.New("Email/Password can't be empty")
	}
	if a.CheckUserExistOrNot(email) {
		return nil, errors.New("user already exists")
	}
	hashedPwd, err := utils.HashPassword(*password)
	if err != nil {
		return nil, errors.New("internal server error happened")
	}

	var user internal.User
	user.Email = *email
	user.Password = hashedPwd
	if err := a.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil

}
