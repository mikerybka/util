package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type TwilioClient struct {
	AccountSID  string
	AuthToken   string
	PhoneNumber string
}

func (c *TwilioClient) SendSMS(to, message string) error {
	apiURL := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", c.AccountSID)

	data := url.Values{}
	data.Set("To", to)
	data.Set("From", c.PhoneNumber)
	data.Set("Body", message)

	req, err := http.NewRequest("POST", apiURL, strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}

	req.SetBasicAuth(c.AccountSID, c.AuthToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send SMS: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)
		return fmt.Errorf("failed to send SMS, status: %d, error: %v", resp.StatusCode, result)
	}

	return nil
}
