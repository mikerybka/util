package util

type Auth struct {
	Accounts      map[string]Account
	SessionTokens map[string]string
}

func (a *Auth) SendLoginCode(phone string) error
func (a *Auth) Login(code string) (string, error)
func (a *Auth) Logout(token string) error
func (a *Auth) WhoAmI(token string) (string, error)
