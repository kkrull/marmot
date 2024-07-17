package coremetarepo

// Maintenance of a meta repository.
type MetaDataAdmin interface {
	Create(metaRepoPath string) error
}
