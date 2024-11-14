package internal

import (
	"encoding/json"
	"os"
)

type PatientRecord struct {
	RecordID  string
	PatientID string
	TestID    string
}

type Storage interface {
	GetAllPatientRecords(PatientID string) []PatientRecord
	GetRecord(RecordID string) *PatientRecord
	InsertNewRecord(PatientID, RecordID, TestID string) error
}

//---

type JSONStorage struct {
	directory string
}

func NewJSONStorage(directory string) *JSONStorage {
	return &JSONStorage{directory}
}

func (s *JSONStorage) GetAllPatientRecords(PatientID string) []PatientRecord {
	files, err := os.ReadDir(s.directory)

	if err != nil {
		return nil
	}

	var records []PatientRecord
	for _, file := range files {
		f, err := os.ReadFile(s.directory + "/" + file.Name())

		if err != nil {
			return nil
		}

		var tempRecord PatientRecord
		json.Unmarshal(f, &tempRecord)

		if tempRecord.PatientID == PatientID {
			records = append(records, tempRecord)
		}
	}

	return records
}

func (s *JSONStorage) GetRecord(RecordID string) *PatientRecord {
	file, err := os.ReadFile(s.directory + "/" + RecordID + ".json")

	if err != nil {
		return nil
	}

	var record *PatientRecord
	json.Unmarshal(file, &record)

	return record
}

func (s *JSONStorage) InsertNewRecord(PatientID, RecordID, TestID string) error {
	return nil
}
