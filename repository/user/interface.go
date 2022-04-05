package user

import _entities "sepatuku-project/entity/user"

type UserRepositoryInterface interface {
	GetUser(id int) (_entities.User, int, error)
	DeleteUser(id int) (_entities.User, error)
	CreateUser(user _entities.User) (_entities.User, error)
	UpdatedUser(users _entities.User) (_entities.User, error)
}
