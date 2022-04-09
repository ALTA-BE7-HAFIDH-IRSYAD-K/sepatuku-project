package user

import (
	"fmt"
	_entitiesproduct "sepatuku-project/entity/product"
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

func (ur *UserRepository) GetUser(id int) (_entities.User, []_entitiesproduct.Product, int, error) {
	var users _entities.User
	var product []_entitiesproduct.Product
	type produk interface{}
	tx := ur.database.Find(&users, id)
	if tx.Error != nil {
		return users, product, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return users, product, 0, tx.Error
	}
	err := ur.database.Where("user_id", id).Find(&product)
	if tx.Error != nil {
		return users, product, 0, err.Error
	}
	return users, product, int(tx.RowsAffected), nil
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
func (ur *UserRepository) CreateUser(user _entities.User) (_entities.User, int, error) {
	if user.Username == "" {
		return user, 0, fmt.Errorf("insert your username")
	}
	if user.Email == "" {
		return user, 0, fmt.Errorf("insert your email")
	}
	if user.Password == "" {
		return user, 0, fmt.Errorf("insert your password")
	}
	if user.Phone == "" {
		return user, 0, fmt.Errorf("insert your phone number")
	}
	tx := ur.database.Save(&user)
	if tx.Error != nil {
		return user, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return user, 2, tx.Error
	}
	return user, int(tx.RowsAffected), nil
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
