package util

type Sender interface {
	Send(to string, msg string) error
}
