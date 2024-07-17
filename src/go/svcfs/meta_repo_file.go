package svcfs

import (
	"encoding/json"
	"fmt"
	"os"
)

// Construct a meta repo file with only minimal information; e.g. no Git repositories.
func EmptyMetaRepoFile(version string) *metaRepoFile {
	return &metaRepoFile{
		MetaRepo: metaRepoData{
			RemoteRepositories: make([]remoteRepositoryData, 0),
		},
		Version: version,
	}
}

// Structure of the meta data file and its representation in JSON.
type metaRepoFile struct {
	MetaRepo metaRepoData `json:"meta_repo"`
	Version  string       `json:"version"`
}

type metaRepoData struct {
	RemoteRepositories []remoteRepositoryData `json:"remote_repositories"`
}

type remoteRepositoryData struct {
	Url string `json:"url"`
}

/* JSON encoding */

func (objectRoot *metaRepoFile) WriteTo(filename string) error {
	var encoder *json.Encoder
	if file, fileErr := os.Create(filename); fileErr != nil {
		return fmt.Errorf("failed to create file %s; %w", filename, fileErr)
	} else {
		defer file.Close()
		encoder = json.NewEncoder(file)
	}

	if encodeErr := encoder.Encode(objectRoot); encodeErr != nil {
		return fmt.Errorf("failed to encode JSON data; %w", encodeErr)
	} else {
		return nil
	}
}
