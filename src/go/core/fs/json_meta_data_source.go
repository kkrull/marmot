package fs

// Stores meta data in JSON files in a directory that Marmot manages
type JsonMetaDataSource struct{}

func (source *JsonMetaDataSource) Init() error {
	return nil
}
