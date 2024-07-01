package util

import (
	"fmt"
)

func NewAuthSystem() *AuthSystem {
	usersTable := NewTable[*User]()
	usersTable.AddUniqConstraint("ID")
	usersTable.AddUniqConstraint("Phone")
	usersTable.AddUniqConstraint("Email")
	return &AuthSystem{
		Users: usersTable,
	}
}

type AuthSystem struct {
	// Config
	TwilioClient    *TwilioClient
	LoginCodeMsgFmt string

	// Data
	Users         *Table[*User]
	LoginCodes    map[string]string // user ID => 6 digit code
	SessionTokens map[string]string // token => user ID
}

func (a *AuthSystem) Register(user *User) error {
	return a.Users.Insert(user)
}

// Creates a new login code.
// Only one login code per user can be active at a time.
func (a *AuthSystem) CreateLoginCode(userID string) string {
	// Generate a 6 digit code.
	code := RandomCode(6)

	// Save the code.
	a.LoginCodes[userID] = code

	return code
}

// CheckLoginCode returns true if the login code is valid.
func (a *AuthSystem) CheckLoginCode(userID, code string) bool {
	if code == "" {
		return false
	}
	return a.LoginCodes[userID] == code
}

// SendLoginCode sends a login code the users phone number on file via Twilio SMS.
func (a *AuthSystem) SendLoginCode(phone string) (userID string, err error) {
	// Find the user.
	users := a.Users.FindBy("Phone", phone)
	userID, user, found := OnlyOne(users)
	if !found {
		return "", fmt.Errorf("user not registered")
	}

	// Create a login code.
	code := RandomCode(6)
	a.LoginCodes[userID] = code

	// Send the login code.
	msg := fmt.Sprintf(a.LoginCodeMsgFmt)
	err = a.TwilioClient.SendSMS(string(user.Phone), msg)
	if err != nil {
		return "", fmt.Errorf("twilio: %s", err)
	}

	// Return the user ID.
	return userID, nil
}

// Login creates a user session.
// It returns a token or error.
func (a *AuthSystem) Login(userID, code string) (token string, err error) {
	ok := a.CheckLoginCode(userID, code)
	if !ok {
		return "", fmt.Errorf("bad login")
	}

	// Generate a unique token
	for {
		token = RandomToken(32)
		if _, ok := a.SessionTokens[token]; !ok {
			break
		}
	}

	// Save the token.
	a.SessionTokens[token] = userID

	return token, nil
}

func (a *AuthSystem) Logout(token string) {
	delete(a.SessionTokens, token)
}

func (a *AuthSystem) WhoAmI(token string) string {
	return a.SessionTokens[token]
}
