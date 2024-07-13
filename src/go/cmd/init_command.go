package cmd

// Initializes a new meta repo where none existed before
type InitCommand struct {
	MetaDataSource MetaDataSource
}

func (command InitCommand) Run() error {
	return command.MetaDataSource.Init()
}

// Access to meta data about Git repositories
type MetaDataSource interface {
	Init() error
}
