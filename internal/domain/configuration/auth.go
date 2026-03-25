package configuration

import (
	"os"
	"strings"
)

func GetDisableAuth() bool {
	serviceToken := os.Getenv("GIT_DISABLE_AUTH")
	return strings.ToLower(serviceToken) == "true"
}
