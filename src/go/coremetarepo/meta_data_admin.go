package coremetarepo

// Maintenance of a meta repository.
type MetaDataAdmin interface {
	Init() error
	InitP(metaRepoPath string) error
}
