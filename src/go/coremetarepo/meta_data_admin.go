package coremetarepo

// Maintenance of a meta repository.
type MetaDataAdmin interface {
	InitP(metaRepoPath string) error
}
