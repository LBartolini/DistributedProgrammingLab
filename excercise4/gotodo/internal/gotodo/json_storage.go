package gotodo

import (
	"encoding/json"
	"os"
)

type JsonStorage struct {
	FileName string
}

func NewJsonStorage(filename string) *JsonStorage {
	return &JsonStorage{filename}
}

func (storage *JsonStorage) Save(data []Todo) error {
	fileData, err := json.MarshalIndent(data, "", "    ")

	if err != nil {
		return err
	}

	return os.WriteFile(storage.FileName, fileData, 0644)
}

func (storage *JsonStorage) Load(data *[]Todo) error {
	fileData, err := os.ReadFile(storage.FileName)

	if err != nil {
		return err
	}

	return json.Unmarshal(fileData, data)
}
