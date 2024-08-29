package testsupportfs

type dirFixtureDown struct {
	prefix string
}

func (state *dirFixtureDown) Create() (dirFixtureState, error) {
	return state, nil
}

func (state *dirFixtureDown) Destroy() (dirFixtureState, error) {
	return state, nil
}
