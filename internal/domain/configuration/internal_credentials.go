package configuration

import (
	"os"
	"strings"
)

func GetDisableAuth() bool {
	disableAuth := os.Getenv("GIT_DISABLE_AUTH")
	return strings.ToLower(disableAuth) == "true"
}
