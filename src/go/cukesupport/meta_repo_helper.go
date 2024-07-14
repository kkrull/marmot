package cukesupport

import "fmt"

var thatMetaRepo string

// Reset the meta repo so another scenario can make its own
func ClearThatMetaRepo() { //TODO KDK: Call this from within SetThatMetaRepo so it's always called?
	thatMetaRepo = ""
}

// A path to a meta repo which has been set earlier in this scenario, or an empty string.
func PeekThatMetaRepo() string {
	return thatMetaRepo
}

func SetThatMetaRepo(path string) string {
	thatMetaRepo = path
	return thatMetaRepo
}

// A path to a meta repo which has been set earlier in this scenario, or an error.
func ThatMetaRepo() (string, error) {
	if thatMetaRepo == "" {
		return "", fmt.Errorf("meta_repo_steps: no meta repo has been configured")
	}

	return thatMetaRepo, nil
}
