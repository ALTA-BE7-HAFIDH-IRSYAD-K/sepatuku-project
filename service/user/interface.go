package user

import _entities "sepatuku-project/entity/user"

type UserServiceInterface interface {
	GetUser(id int) (_entities.User, int, error)
	DeleteUser(id int) (_entities.User, error)
	CreateUser(user _entities.User) (_entities.User, error)
	UpdatedUser(users _entities.User, id int) (_entities.User, error)
}
