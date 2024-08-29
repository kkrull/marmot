package testsupportfs

// Make a fixture that manages a temporary directory to use for testing.
func NewDirFixture(prefix string) *DirFixture {
	return &DirFixture{
		state: &dirFixtureDown{prefix: prefix},
	}
}

// Creates a temporary test directory and deletes it when done testing.
type DirFixture struct {
	state dirFixtureState
}

// Ensure the test directory has been created; re-entrant.
func (fixture *DirFixture) Setup() error {
	up, createErr := fixture.state.Create()
	fixture.state = up
	return createErr
}

// Ensure any test directory that was created before has been deleted; re-entrant.
func (fixture *DirFixture) Teardown() error {
	down, destroyErr := fixture.state.Destroy()
	fixture.state = down
	return destroyErr
}

type dirFixtureState interface {
	Create() (dirFixtureState, error)
	Destroy() (dirFixtureState, error)
}
