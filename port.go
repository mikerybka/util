package util

func Port() string {
	port := EnvVar("PORT", "3000")
	return ":" + port
}
