package cukesupport

import "fmt"

var thatMetaRepo string

// A path to a meta repo which has been set earlier in this scenario, or an empty string.
func PeekThatMetaRepo() string {
	return thatMetaRepo
}

// A path to a meta repo which has been set earlier in this scenario, or an error.
func ThatMetaRepo() (string, error) {
	if thatMetaRepo == "" {
		return "", fmt.Errorf("meta_repo_steps: no meta repo has been configured")
	}

	return thatMetaRepo, nil
}

func SetThatMetaRepo(path string) string {
	thatMetaRepo = path
	return thatMetaRepo
}
