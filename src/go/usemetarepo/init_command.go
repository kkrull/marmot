package usemetarepo

// Initializes a new meta repo where none existed before
type InitCommand struct {
	MetaDataSource MetaDataSource
}

func (command InitCommand) Run() error {
	return command.MetaDataSource.Init()
}
