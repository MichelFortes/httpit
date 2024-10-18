package cli

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/MichelFortes/httpit/internal/constraints"
	"github.com/MichelFortes/httpit/pkg/model"
)

func GetTestScheme() (*model.TestScheme, error) {

	scheme := &model.TestScheme{}

	args := os.Args
	if len(args) != 2 {
		return scheme, errors.New(constraints.ErrorMissingConfigArg)
	}

	file, err := os.Open(args[1])
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
