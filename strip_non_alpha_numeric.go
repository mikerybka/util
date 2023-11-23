package util

func StripNonAlphaNumeric(s string) string {
	stripped := ""
	for _, ch := range s {
		if IsAlphaNumeric(ch) {
			stripped += string(ch)
		}
	}
	return stripped
}
