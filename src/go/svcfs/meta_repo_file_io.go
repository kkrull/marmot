package svcfs

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadLocalMetaRepoFile(filename string) (*localRootObjectData, error) {
	root := &localRootObjectData{}
	if err := readMetaRepoFile(filename, root); err != nil {
		return nil, err
	} else {
		return root, nil
	}
}

func ReadSharedMetaRepoFile(filename string) (*sharedRootObjectData, error) {
	root := &sharedRootObjectData{}
	if err := readMetaRepoFile(filename, root); err != nil {
		return nil, err
	} else {
		return root, nil
	}
}

func readMetaRepoFile(filename string, root any) error {
	var decoder *json.Decoder
	if file, openErr := os.Open(filename); openErr != nil {
		return fmt.Errorf("failed to open file %s; %w", filename, openErr)
	} else {
		defer file.Close()
		decoder = json.NewDecoder(file)
	}

	if decodeErr := decoder.Decode(&root); decodeErr != nil {
		return fmt.Errorf("failed to decode %s; %w", filename, decodeErr)
	} else {
		return nil
	}
}

func (root *localRootObjectData) WriteTo(filename string) error {
	var encoder *json.Encoder
	if file, fileErr := os.Create(filename); fileErr != nil {
		return fmt.Errorf("failed to create file %s; %w", filename, fileErr)
	} else {
		defer file.Close()
		encoder = json.NewEncoder(file)
		encoder.SetIndent("", "  ")
	}

	if encodeErr := encoder.Encode(root); encodeErr != nil {
		return fmt.Errorf("failed to encode JSON data; %w", encodeErr)
	} else {
		return nil
	}
}

func (root *sharedRootObjectData) WriteTo(filename string) error {
	var encoder *json.Encoder
	if file, fileErr := os.Create(filename); fileErr != nil {
		return fmt.Errorf("failed to create file %s; %w", filename, fileErr)
	} else {
		defer file.Close()
		encoder = json.NewEncoder(file)
		encoder.SetIndent("", "  ")
	}

	if encodeErr := encoder.Encode(root); encodeErr != nil {
		return fmt.Errorf("failed to encode JSON data; %w", encodeErr)
	} else {
		return nil
	}
}
