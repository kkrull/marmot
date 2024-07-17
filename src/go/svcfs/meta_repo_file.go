package svcfs

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kkrull/marmot/corerepository"
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

/* Core conversion */

func (objectRoot *metaRepoFile) ToCoreRepositories() (corerepository.Repositories, error) {
	repositories := make([]corerepository.Repository, len(objectRoot.MetaRepo.RemoteRepositories))
	for i, remoteRepositoryData := range objectRoot.MetaRepo.RemoteRepositories {
		if remoteRepository, parseErr := corerepository.RemoteRepositoryS(remoteRepositoryData.Url); parseErr != nil {
			return nil, fmt.Errorf("failed to parse %s; %w", remoteRepositoryData.Url, parseErr)
		} else {
			repositories[i] = remoteRepository
		}
	}

	return &corerepository.RepositoriesArray{
		Repositories: repositories,
	}, nil
}

/* JSON decoding */

func ReadMetaRepoFile(filename string) (*metaRepoFile, error) {
	var decoder *json.Decoder
	if file, openErr := os.Open(filename); openErr != nil {
		return nil, fmt.Errorf("failed to open file %s; %w", filename, openErr)
	} else {
		// defer metaDataFd.Close() //TODO KDK: Test and restore
		decoder = json.NewDecoder(file)
	}

	var content metaRepoFile
	if decodeErr := decoder.Decode(&content); decodeErr != nil {
		return nil, fmt.Errorf("failed to decode %s; %w", filename, decodeErr)
	} else {
		return &content, nil
	}
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
