package greeting

import "fmt"

func Message() string {
	return "Hello from Go Releases! 🚀"
}

func BuildInfo(version, goos, arch string) string {
	return fmt.Sprintf("Version: %s\nOS: %s\nArch: %s", version, goos, arch)
}
