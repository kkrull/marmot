package testsupportfs

import "os"

type dirFixtureDown struct {
	prefix string
}

func (down *dirFixtureDown) Create() (dirFixtureState, error) {
	if tempDir, mkdirErr := os.MkdirTemp("", down.prefix); mkdirErr != nil {
		return down, mkdirErr
	} else {
		return &dirFixtureUp{path: tempDir, prefix: down.prefix}, nil
	}
}

func (down *dirFixtureDown) Destroy() (dirFixtureState, error) {
	return down, nil
}
