package auth

type AuthRepositoryInterface interface {
	Login(identifier string, password string) (string, error)
}
