package coremetarepo

// Manages the life cycle a meta repository.
type MetaDataAdmin interface {
	// Create a meta repository at the specified path on the local filesystem.
	Create(metaRepoPath string) error

	// Returns true if the specified path exists and is already a meta repository.
	IsMetaRepo(metaRepoPath string) bool
}
