package config

import (
	"encoding/json"
	"os"

	"github.com/MichelFortes/httpit/pkg/model"
)

func GetTestScheme(filename string) (*model.TestScheme, error) {

	scheme := &model.TestScheme{}

	file, err := os.Open(filename)
	if err != nil {
		return scheme, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(scheme)
	if err != nil {
		return scheme, err
	}

	return scheme, nil
}
