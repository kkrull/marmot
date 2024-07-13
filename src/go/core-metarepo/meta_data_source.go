package core_metarepo

// Access to meta data about Git repositories
type MetaDataSource interface {
	Init() error
}
