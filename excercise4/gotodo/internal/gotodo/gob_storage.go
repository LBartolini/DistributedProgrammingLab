package gotodo

import (
	"encoding/gob"
	"os"
)

type GOBStorage struct {
	FileName string
}

func NewGOBStorage(filename string) *GOBStorage {
	return &GOBStorage{filename}
}

func (storage *GOBStorage) Save(data []Todo) error {
	file, err := os.OpenFile(storage.FileName, os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		return err
	}

	enc := gob.NewEncoder(file)

	return enc.Encode(data)
}

func (storage *GOBStorage) Load(data *[]Todo) error {
	file, err := os.Open(storage.FileName)

	if err != nil {
		return err
	}

	enc := gob.NewDecoder(file)

	return enc.Decode(data)
}
