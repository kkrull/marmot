package svcfs

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadMetaRepoFile(filename string) (*rootObjectData, error) {
	var decoder *json.Decoder
	if file, openErr := os.Open(filename); openErr != nil {
		return nil, fmt.Errorf("failed to open file %s; %w", filename, openErr)
	} else {
		// defer metaDataFd.Close() //TODO KDK: Test and restore
		decoder = json.NewDecoder(file)
	}

	var root rootObjectData
	if decodeErr := decoder.Decode(&root); decodeErr != nil {
		return nil, fmt.Errorf("failed to decode %s; %w", filename, decodeErr)
	} else {
		return &root, nil
	}
}

func (root *rootObjectData) WriteTo(filename string) error {
	var encoder *json.Encoder
	if file, fileErr := os.Create(filename); fileErr != nil {
		return fmt.Errorf("failed to create file %s; %w", filename, fileErr)
	} else {
		defer file.Close()
		encoder = json.NewEncoder(file)
	}

	if encodeErr := encoder.Encode(root); encodeErr != nil {
		return fmt.Errorf("failed to encode JSON data; %w", encodeErr)
	} else {
		return nil
	}
}
