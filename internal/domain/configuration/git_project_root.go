package configuration

import "os"

// GetGitProjectRoot returns the git project root from environment variable or default
func GetGitProjectRoot() string {
	gitProjectRoot := os.Getenv("GIT_PROJECT_ROOT")
	if gitProjectRoot == "" {
		gitProjectRoot = "/data/repos"
	}
	return gitProjectRoot
}
