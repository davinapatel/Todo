package main

import (
	"encoding/json"
	"os"
)

// [] After name of struct, means using generics
// any is an interface
type Storage[T any] struct {
	FileName string
}

func NewStorage[T any](filename string) *Storage[T] {
	return &Storage[T]{FileName: filename}
}

func (s *Storage[T]) Save(data T) error {
	// convert data into json
	// (data, prefix, indentation for nice formatting of json)
	fileData, err := json.MarshalIndent(data, "", "    ")

	if err != nil {
		return err
	}

	// write json to file
	// 0644 = permission type, owner can r&w, everyone else just r
	return os.WriteFile(s.FileName, fileData, 0644)
}

func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.FileName)

	if err != nil {
		return err
	}

	// Convert json data and populate the data arg
	return json.Unmarshal(fileData, data)
}
