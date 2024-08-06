package coremetarepo

// Manages the life cycle a meta repository.
type MetaDataAdmin interface {
	// Create a meta repository at the specified path on the local filesystem.
	Create(metaRepoPath string) error
}
