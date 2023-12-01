package util

func PackageManager() string {
	os := OS()
	switch os {
	case "darwin":
		return "homebrew"
	case "debian", "ubuntu":
		return "apt"
	case "arch", "manjaro":
		return "pacman"
	case "fedora", "rhel":
		return "dnf"
	default:
		return os
	}
}
