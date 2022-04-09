package user

import (
	_entitiesproduct "sepatuku-project/entity/product"
	_entities "sepatuku-project/entity/user"
)

type UserRepositoryInterface interface {
	GetUser(id int) (_entities.User, []_entitiesproduct.Product, int, error)
	DeleteUser(id int) (_entities.User, error)
	CreateUser(user _entities.User) (_entities.User, int, error)
	UpdatedUser(user _entities.User) (_entities.User, error)
}
