package cmd

type InitCmd struct {
	MetaDataSource MetaDataSource
}

func (command InitCmd) Run() error {
	command.MetaDataSource.EnsureCreated()
	return nil
}

type MetaDataSource interface {
	EnsureCreated()
}
