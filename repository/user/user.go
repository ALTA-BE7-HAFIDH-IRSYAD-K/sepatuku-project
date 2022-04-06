package user

import (
	_entities "sepatuku-project/entity/user"

	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

func (ur *UserRepository) GetUser(id int) (_entities.User, int, error) {
	var user _entities.User
	tx := ur.database.Find(&user, id)
	if tx.Error != nil {
		return user, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return user, 0, tx.Error
	}
	return user, int(tx.RowsAffected), nil
}

func (ur *UserRepository) DeleteUser(id int) (_entities.User, error) {
	var user _entities.User
	tx := ur.database.Delete(&user, id)
	if tx.Error != nil {
		return user, tx.Error
	}
	if tx.RowsAffected == 0 {
		return user, tx.Error

	}
	return user, nil
}
func (ur *UserRepository) CreateUser(user _entities.User) (_entities.User, error) {
	tx := ur.database.Save(&user)
	if tx.Error != nil {
		return user, tx.Error
	}
	if tx.RowsAffected == 0 {
		return user, tx.Error

	}
	return user, nil
}
func (ur *UserRepository) UpdatedUser(user _entities.User) (_entities.User, error) {
	tx := ur.database.Save(&user)
	if tx.Error != nil {
		return user, tx.Error
	}
	if tx.RowsAffected == 0 {
		return user, tx.Error
	}
	return user, nil
}
