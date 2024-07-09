package cmd

type InitCmd struct {
	MetaDataStore MetaDataStore
}

func (command InitCmd) Run() error {
	command.MetaDataStore.EnsureCreated()
	return nil
}

type MetaDataStore interface {
	EnsureCreated()
}
