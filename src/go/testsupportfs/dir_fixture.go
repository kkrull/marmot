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

/* Life cycle */

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
	BasePath() (string, error)
	Create() (dirFixtureState, error)
	Destroy() (dirFixtureState, error)
	PathBuilder() (*PathBuilder, error)
}

/* Use */

// Get the path to the test directory, if it has been created.
func (fixture *DirFixture) BasePath() (string, error) {
	return fixture.state.BasePath()
}

// Get something to build paths inside of the test directory, provided it has been created.
func (fixture *DirFixture) PathBuilder() (*PathBuilder, error) {
	return fixture.state.PathBuilder()
}
