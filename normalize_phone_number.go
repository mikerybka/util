package util

func NormalizePhoneNumber(number string) string {
	number = FilterString(number, Digits)
	if len(number) == 11 && number[0] == 1 {
		number = number[1:]
	}
	return number
}
