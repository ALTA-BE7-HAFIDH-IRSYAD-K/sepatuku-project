package auth

type AuthServiceInterface interface {
	Login(identifier string, password string) (string, error)
}
