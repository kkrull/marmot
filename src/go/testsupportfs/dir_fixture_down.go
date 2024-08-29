package testsupportfs

import (
	"fmt"
	"os"
)

type dirFixtureDown struct {
	prefix string
}

func (down *dirFixtureDown) BasePath() (string, error) {
	return "", fmt.Errorf("test directory starting with <%s> not created yet", down.prefix)
}

func (down *dirFixtureDown) Create() (dirFixtureState, error) {
	if tempDir, mkdirErr := os.MkdirTemp("", down.prefix); mkdirErr != nil {
		return down, mkdirErr
	} else {
		return &dirFixtureUp{
			createdWithPrefix: down.prefix,
			dirPath:           tempDir,
		}, nil
	}
}

func (down *dirFixtureDown) Destroy() (dirFixtureState, error) {
	return down, nil
}
