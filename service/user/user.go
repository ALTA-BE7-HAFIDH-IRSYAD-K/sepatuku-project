package user

import (
	"sepatuku-project/delivery/response"
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
	user, rows, err := uuc.userRepository.GetUser(id)
	return user, rows, err
}

func (uuc *UserService) DeleteUser(id int) (_entities.User, error) {
	user, err := uuc.userRepository.DeleteUser(id)
	return user, err
}

func (uuc *UserService) CreateUser(user _entities.User) (_entities.User, error) {
	user, err := uuc.userRepository.CreateUser(user)
	password, err := response.HashPassword(user.Password)
	user.Password = password
	return user, err
}

func (uuc *UserService) UpdatedUser(user _entities.User, id int) (_entities.User, error) {
	user.ID = uint(id)
	user, err := uuc.userRepository.UpdatedUser(user)
	password, err := response.HashPassword(user.Password)
	user.Password = password
	return user, err
}
