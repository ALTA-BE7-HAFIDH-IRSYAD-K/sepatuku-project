package user

import (
	"sepatuku-project/delivery/response"
	_entitiesproduct "sepatuku-project/entity/product"
	_entities "sepatuku-project/entity/user"
	_userRepository "sepatuku-project/repository/user"

	"github.com/jinzhu/copier"
)

type UserService struct {
	userRepository _userRepository.UserRepositoryInterface
}

func NewUserService(userRepo _userRepository.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		userRepository: userRepo,
	}
}

func (uuc *UserService) GetUser(id int) (_entities.User, []_entitiesproduct.Product, int, error) {
	user, product, rows, err := uuc.userRepository.GetUser(id)
	return user, product, rows, err
}
func (uuc *UserService) GetProfile(id int) (_entities.UserRespon, []_entitiesproduct.Product, int, error) {
	user, product, rows, err := uuc.userRepository.GetUser(id)
	userRes := _entities.UserRespon{}
	copier.Copy(&userRes, &user)
	return userRes, product, rows, err
}

func (uuc *UserService) DeleteUser(id int) (_entities.User, int, error) {
	user, rowDel, err := uuc.userRepository.DeleteUser(id)
	return user, rowDel, err
}

func (uuc *UserService) CreateUser(user _entities.User) (_entities.UserRespon, int, error) {
	password, row, err := response.HashPassword(user.Password)
	user.Password = password
	user, row, err = uuc.userRepository.CreateUser(user)
	userRes := _entities.UserRespon{}
	copier.Copy(&userRes, &user)
	return userRes, row, err
}

func (uuc *UserService) UpdatedUser(user _entities.User, id int) (_entities.UserRespon, error) {
	user.ID = uint(id)
	password, _, err := response.HashPassword(user.Password)
	user.Password = password
	user, err = uuc.userRepository.UpdatedUser(user)
	userRes := _entities.UserRespon{}
	copier.Copy(&userRes, &user)
	return userRes, err
}
