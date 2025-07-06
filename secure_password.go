package util

var capitals, lowercases, digits = "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "abcdefghijklmnopqrstuvwxyz", "1234567890"

func SecurePassword() string {
	return RandomString(12, capitals+lowercases+digits)
}
