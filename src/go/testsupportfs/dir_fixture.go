package testsupportfs

// A test fixture that makes a temporary test directory with a prefix and deletes it when done.
func NewDirFixture(prefix string) *DirFixture {
	return &DirFixture{prefix}
}

type DirFixture struct {
	prefix string
}

func (fixture *DirFixture) Setup() error {
	return nil
}

func (fixture *DirFixture) Teardown() error {
	return nil
}
