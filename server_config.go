package util

type ServerConfig struct {
	AdminPhone string
	Twilio     *TwilioClient
	DataDir    string
	CodeDir    string
}
