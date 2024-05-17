package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type AuthHandler struct {
	AuthFiles FileSystem
	Twilio    *TwilioClient
}

func (a *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/send-login-code" {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		input := struct {
			Phone string
		}{}
		err = json.Unmarshal(b, &input)
		if err != nil {
			panic(err)
		}
		err = a.SendLoginCode(input.Phone)
		if err != nil {
			panic(err)
		}
		return
	}

	if r.URL.Path == "/login" {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		input := struct {
			Phone string
			Code  string
		}{}
		err = json.Unmarshal(b, &input)
		if err != nil {
			panic(err)
		}
		session, err := a.Login(input.Phone, input.Code)
		if err != nil {
			panic(err)
		}
		json.NewEncoder(w).Encode(session)
		return
	}
}

func (a *AuthHandler) SendLoginCode(phone string) error {
	// Reject malformed input.
	if !IsTenDigits(phone) {
		return fmt.Errorf("expected 10 digit phone number")
	}

	// Generate a random code for the user.
	code := RandomCode(6)

	// Write the code to file.
	loginCodePath := fmt.Sprintf("/users/%s/login-codes/%s", phone, code)
	err := a.AuthFiles.WriteFile(loginCodePath, []byte(code))
	if err != nil {
		panic(err)
	}

	// Send the code to the phone number.
	msg := fmt.Sprintf("Your login code is %s", code)
	err = a.Twilio.SendSMS(phone, msg)
	if err != nil {
		panic(err)
	}

	// Wait 5 minutes in another goroutine, then delete the login code.
	go func() {
		time.Sleep(5 * time.Minute)
		a.AuthFiles.Remove(loginCodePath)
	}()

	return nil
}

func (a *AuthHandler) Login(phone, code string) (*Session, error) {
	// Check if the login code is valid.
	loginCodePath := fmt.Sprintf("/users/%s/login-codes/%s", phone, code)
	ok := a.AuthFiles.IsFile(loginCodePath)
	if !ok {
		return nil, fmt.Errorf("bad code")
	}

	// If so, delete the code.
	err := a.AuthFiles.Remove(loginCodePath)
	if err != nil {
		panic(err)
	}

	// And create a new session.
	var token, sessionPath string
	for {
		token = RandomToken(8)
		sessionPath = fmt.Sprintf("/users/%s/sessions/%s", phone, token)
		if !a.AuthFiles.IsFile(sessionPath) {
			break
		}
	}
	s := Session{
		UserID: phone,
		Token:  token,
	}
	b, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	err = a.AuthFiles.WriteFile(sessionPath, b)
	if err != nil {
		panic(err)
	}

	return &s, nil
}
