package testsupportdata

import (
	"fmt"
	"os"

	expect "github.com/kkrull/marmot/testsupportexpect"
)

// Create a fixture that manages a local file system rooted at a path with the given prefix.
func NewPathBuilder(pathPrefix string) *PathBuilder {
	return &PathBuilder{pathPrefix: pathPrefix}
}

// Builds paths on this filesystem, according to specifications.
type PathBuilder struct {
	originalCwd string
	pathPrefix  string
	testFsRoot  string
}

/* Builder */

func (builder *PathBuilder) Build() (string, error) {
	return "/home/user/bogus", nil
}

func (builder *PathBuilder) Missing() *PathBuilder {
	return builder
}

/* Fixture */

// TODO KDK: Extract states for the directory existing+set or not
func (builder *PathBuilder) Setup() error {
	testFsRoot = expect.NoError(os.MkdirTemp("", "RegisterLocalRepositoriesCommand-"))
	originalCwd = expect.NoError(os.Getwd())
	Expect(os.Chdir(testFsRoot)).To(Succeed())
	return nil
}

func (fixture *PathBuilder) Teardown() error {
	if chdirErr := os.Chdir(fixture.originalCwd); chdirErr != nil {
		return fmt.Errorf("failed to pop back to directory %s; %w", fixture.originalCwd, chdirErr)
	} else if removeErr := os.RemoveAll(fixture.testFsRoot); removeErr != nil {
		return fmt.Errorf("failed to remove directory %s; %w", fixture.testFsRoot, removeErr)
	} else {
		return nil
	}
}
