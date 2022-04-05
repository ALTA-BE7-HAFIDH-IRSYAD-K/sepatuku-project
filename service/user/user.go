package user

import (
	_entities "sepatuku-project/entity/user"
	_userRepository "sepatuku-project/repository/user"
)

type UserService struct {
	userRepository _userRepository.UserRepositoryInterface
}

func NewUserService(userRepo _userRepository.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		userRepository: userRepo,
	}
}

func (uuc *UserService) GetUser(id int) (_entities.User, int, error) {
	users, rows, err := uuc.userRepository.GetUser(id)
	return users, rows, err
}

func (uuc *UserService) DeleteUser(id int) (_entities.User, error) {
	users, err := uuc.userRepository.DeleteUser(id)
	return users, err
}
func (uuc *UserService) CreateUser(user _entities.User) (_entities.User, error) {
	users, err := uuc.userRepository.CreateUser(user)
	return users, err
}
func (uuc *UserService) UpdatedUser(users _entities.User, id int) (_entities.User, error) {
	users.ID = uint(id)
	users, err := uuc.userRepository.UpdatedUser(users)
	return users, err
}
