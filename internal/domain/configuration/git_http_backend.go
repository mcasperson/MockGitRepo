package configuration

import "os"

const (
	// This is /usr/lib/git-core/git-http-backend on ubuntu
	// It is this on a mac
	// /Library/Developer/CommandLineTools/usr/libexec/git-core/git-http-backend
	gitHTTPBackendPath = "/usr/libexec/git-core/git-http-backend"
)

func GetGitHttpBackend() string {
	gitHttpBackend := os.Getenv("GIT_HTTP_BACKEND")
	if gitHttpBackend == "" {
		gitHttpBackend = gitHTTPBackendPath
	}
	return gitHttpBackend
}
