package testsupportfs

// Builds paths with certain properties inside of an existing base directory.
type PathBuilder struct {
	basePath string
}

func (builder *PathBuilder) AsAbsolute() string {
	return builder.basePath
}
