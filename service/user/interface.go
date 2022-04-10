package user

import (
	_entitiesproduct "sepatuku-project/entity/product"
	_entities "sepatuku-project/entity/user"
)

type UserServiceInterface interface {
	GetUser(id int) (_entities.User, []_entitiesproduct.Product, int, error)
	GetProfile(id int) (_entities.UserRespon, []_entitiesproduct.Product, int, error)
	DeleteUser(id int) (_entities.User, int, error)
	CreateUser(user _entities.User) (_entities.UserRespon, int, error)
	UpdatedUser(user _entities.User, id int) (_entities.UserRespon, error)
}
