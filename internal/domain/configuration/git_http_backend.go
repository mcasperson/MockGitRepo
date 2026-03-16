package configuration

import "os"

const (
	// This is /usr/lib/git-core/git-http-backend on ubuntu
	gitHTTPBackendPath = "/usr/libexec/git-core/git-http-backend"
)

func GetGitHttpBackend() string {
	gitHttpBackend := os.Getenv("GIT_HTTP_BACKEND")
	if gitHttpBackend == "" {
		gitHttpBackend = gitHTTPBackendPath
	}
	return gitHttpBackend
}
