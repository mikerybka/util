package util

type WebApp struct {
	Name             string
	Description      string
	Author           string
	Keywords         []string
	Favicon          []byte
	Icon             []byte
	Types            map[string]Type
	CoreResourceType string
	TwilioClient     *TwilioClient
	Files            FileSystem
}
