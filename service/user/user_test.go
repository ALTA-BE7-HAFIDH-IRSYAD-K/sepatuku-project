package user

import (
	"fmt"
	"sepatuku-project/entity/product"
	"sepatuku-project/entity/user"
	_entities "sepatuku-project/entity/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	t.Run("TestGetUserSuccess", func(t *testing.T) {
		userService := NewUserService(mockUserRepository{})
		users, product, rows, err := userService.GetUser(1)
		assert.Nil(t, err)
		assert.Equal(t, "nasrul", users.Username)
		assert.Equal(t, "jordan", product[0].Name_product)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestGetUserError", func(t *testing.T) {
		userService := NewUserService(mockUserRepositoryError{})
		users, products, rows, err := userService.GetUser(1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, rows)
		assert.Equal(t, user.User{}, users)
		assert.Equal(t, []product.Product{}, products)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("TestDeleteUserSuccess", func(t *testing.T) {
		userService := NewUserService(mockUserRepository{})
		users, row, err := userService.DeleteUser(1)
		assert.Nil(t, err)
		assert.Equal(t, "deleted", users.Username)
		assert.Equal(t, 1, row)
	})

	t.Run("TestDeleteUserError", func(t *testing.T) {
		userService := NewUserService(mockUserRepositoryError{})
		users, row, err := userService.DeleteUser(1)
		assert.NotNil(t, err)
		assert.Equal(t, "nasrul", users.Username)
		assert.Equal(t, 0, row)
	})
}

func TestCreateUser(t *testing.T) {
	t.Run("TestCreateUserSuccess", func(t *testing.T) {
		userService := NewUserService(mockUserRepository{})
		users, row, err := userService.CreateUser(user.User{
			Username: "nasrul", Email: "nasrul@mail.com", Password: "12345"})
		assert.Nil(t, err)
		assert.Equal(t, "nasrul", users.Username)
		assert.Equal(t, "nasrul@mail.com", users.Email)
		assert.Equal(t, 1, row)
	})

	t.Run("TestCreateUserError", func(t *testing.T) {
		userService := NewUserService(mockUserRepositoryError{})
		users, row, err := userService.CreateUser(user.User{
			Username: "nasrul", Email: "nasrul@mail.com", Password: "12345"})
		assert.NotNil(t, err)
		assert.Equal(t, _entities.User{}, users)
		assert.Equal(t, 0, row)
	})
}

func TestUpdatedUser(t *testing.T) {
	t.Run("TestUpdatedUserSuccess", func(t *testing.T) {
		userService := NewUserService(mockUserRepository{})
		users, err := userService.UpdatedUser(user.User{
			Username: "nasrullah", Email: "nasrullah@mail.com"}, 1)
		assert.Nil(t, err)
		assert.Equal(t, "nasrullah", users.Username)
		assert.Equal(t, "nasrullah@mail.com", users.Email)
	})

	t.Run("TestUpdatedUserError", func(t *testing.T) {
		userService := NewUserService(mockUserRepositoryError{})
		users, err := userService.UpdatedUser(user.User{
			Username: "nasrullah", Email: "nasrullah@mail.com"}, 1)
		assert.NotNil(t, err)
		assert.Equal(t, user.User{Username: "nasrul", Email: "nasrul@mail.com"}, users)
	})
}

// mock succes
type mockUserRepository struct{}

func (m mockUserRepository) GetUser(id int) (_entities.User, []product.Product, int, error) {
	return _entities.User{
			Username: "nasrul", Email: "nasrul@mail.com"},
		[]product.Product{{Name_product: "jordan"}}, 1, nil
}

func (m mockUserRepository) DeleteUser(id int) (_entities.User, int, error) {
	return _entities.User{
		Username: "deleted", Email: "deleted", Password: "deleted",
	}, 1, nil
}

func (m mockUserRepository) UpdatedUser(user _entities.User) (_entities.User, error) {
	return _entities.User{
			Username: "nasrullah", Email: "nasrullah@mail.com"},
		nil
}
func (m mockUserRepository) CreateUser(user _entities.User) (_entities.User, int, error) {
	return _entities.User{
		Username: "nasrul", Email: "nasrul@mail.com"}, 1, nil
}

//  mock error

type mockUserRepositoryError struct{}

func (m mockUserRepositoryError) GetUser(id int) (_entities.User, []product.Product, int, error) {
	return _entities.User{}, []product.Product{}, 0, fmt.Errorf("error get user")
}
func (m mockUserRepositoryError) DeleteUser(id int) (_entities.User, int, error) {
	return _entities.User{
		Username: "nasrul", Email: "nasrul@mail.com",
	}, 0, fmt.Errorf("error delete user")
}
func (m mockUserRepositoryError) CreateUser(_entities.User) (_entities.User, int, error) {
	return _entities.User{}, 0, fmt.Errorf("error create user")
}
func (m mockUserRepositoryError) UpdatedUser(_entities.User) (_entities.User, error) {
	return _entities.User{
		Username: "nasrul", Email: "nasrul@mail.com", Password: "12345",
	}, fmt.Errorf("error updated user")
}
