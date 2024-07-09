package cmd

type InitCmd struct {
	FileSystem MarmotFileSystem
}

func (command InitCmd) Run() error {
	command.FileSystem.EnsureExists("/path/to/meta")
	return nil
}

type MarmotFileSystem interface {
	EnsureExists(path string)
}
