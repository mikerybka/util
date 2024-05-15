package util

type AuthMethod interface {
	SendLoginCode(phone string) error
	Login(code string) (string, error)
	Logout(token string) error
	WhoAmI(token string) (string, error)
}
