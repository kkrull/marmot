package cmd

// Initializes a new meta repo where none existed before
type InitCmd struct {
	MetaDataSource MetaDataSource
}

func (command InitCmd) Run() error {
	return command.MetaDataSource.Init()
}

// Access to meta data about Git repositories
type MetaDataSource interface {
	Init() error
}
