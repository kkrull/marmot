package coremetarepo

// Maintenance of a meta repository.
type MetaDataAdmin interface {
	Init(metaRepoPath string) error
}
