package auth

import _authRepository "sepatuku-project/repository/auth"

type AuthService struct {
	authRepository _authRepository.AuthRepositoryInterface
}

func NewAuthService(authRepo _authRepository.AuthRepositoryInterface) AuthServiceInterface {
	return &AuthService{
		authRepository: authRepo,
	}
}
func (auc *AuthService) Login(identifier string, password string) (string, error) {
	token, err := auc.authRepository.Login(identifier, password)
	return token, err
}
