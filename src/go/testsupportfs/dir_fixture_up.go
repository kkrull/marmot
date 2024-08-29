package testsupportfs

import "os"

type dirFixtureUp struct {
	path   string
	prefix string
}

func (up *dirFixtureUp) Create() (dirFixtureState, error) {
	return up, nil
}

func (up *dirFixtureUp) Destroy() (dirFixtureState, error) {
	if rmErr := os.RemoveAll(up.path); rmErr != nil {
		return up, rmErr
	}

	return &dirFixtureDown{prefix: up.prefix}, nil
}
