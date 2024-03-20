package utils

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Rezapahlevi3108/hacktiv-assignment-3/internal/model"
)

func WriteStatusToFile(status model.Status, filePath string) error {
	data, err := json.Marshal(status)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
