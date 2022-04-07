package auth

import (
	"errors"
	_middlewares "sepatuku-project/delivery/middleware"
	"sepatuku-project/delivery/response"
	_entities "sepatuku-project/entity/user"

	"gorm.io/gorm"
)

type AuthRepository struct {
	database *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		database: db,
	}
}

func (ar *AuthRepository) Login(identifier string, password string) (string, error) {
	var user _entities.User
	tx := ar.database.Where("email = ? OR username = ?", identifier, identifier).Find(&user)
	if tx.Error != nil {
		return "failed", tx.Error
	}
	if tx.RowsAffected == 0 {
		return "user not found", errors.New("user not found")
	}
	if response.CheckPasswordHash(password, user.Password) != true {
		return "password incorrect", errors.New("password incorrect")
	}
	token, err := _middlewares.CreateToken(int(user.ID), user.Username)
	if err != nil {
		return "create token failed", err
	}
	return token, nil
}
