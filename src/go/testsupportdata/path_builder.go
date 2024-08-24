package testsupportdata

import (
	"errors"
	"os"
	"path/filepath"

	. "github.com/onsi/gomega"
)

func NewPathBuilder(testFsRoot string) *PathBuilder {
	return &PathBuilder{testFsRoot: testFsRoot}
}

// Builds paths on this filesystem, according to specifications.
type PathBuilder struct {
	testFsRoot string
}

func (builder *PathBuilder) Build() string {
	return builder.testFsRoot
}

func (builder *PathBuilder) MissingPath() string {
	path := filepath.Join(builder.testFsRoot, "missing")
	_, statErr := os.Stat(path)
	Expect(errors.Is(statErr, os.ErrNotExist)).To(BeTrue())
	return path
}
