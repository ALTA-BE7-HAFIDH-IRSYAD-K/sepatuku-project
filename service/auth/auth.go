package auth

import (
	"sepatuku-project/delivery/response"
	_entities "sepatuku-project/entity/user"
	_authRepository "sepatuku-project/repository/auth"
)

type AuthService struct {
	authRepository _authRepository.AuthRepositoryInterface
}

func NewAuthService(authRepo _authRepository.AuthRepositoryInterface) AuthServiceInterface {
	return &AuthService{
		authRepository: authRepo,
	}
}
func (auc *AuthService) Login(identifier string, password string) (string, error) {
	var user _entities.User
	token, err := auc.authRepository.Login(identifier, password)
	checkHash, _ := response.CheckPasswordHash(password, user.Password)
	if err != nil && !checkHash {
		return token, err
	}
	return token, err
}
