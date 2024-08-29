package testsupportfs

import (
	"os"
)

type dirFixtureUp struct {
	createdWithPrefix string
	dirPath           string
}

func (up *dirFixtureUp) BasePath() (string, error) {
	return up.dirPath, nil
}

func (up *dirFixtureUp) Create() (dirFixtureState, error) {
	return up, nil
}

func (up *dirFixtureUp) Destroy() (dirFixtureState, error) {
	if rmErr := os.RemoveAll(up.dirPath); rmErr != nil {
		return up, rmErr
	}

	return &dirFixtureDown{prefix: up.createdWithPrefix}, nil
}
