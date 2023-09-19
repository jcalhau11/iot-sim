package oven

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type File struct {
	Path string
}

func (file *File) ReadFile() (any, error) {
	fl, flErr := os.Open(file.Path)
	if flErr != nil {
		return nil, errors.New("File not found")
	}

	defer fl.Close()

	var fileContent interface{}
	decoder := json.NewDecoder(fl)
	if dError := decoder.Decode(&fileContent); dError != nil {
		fmt.Println(dError)
		return nil, errors.New("Could not decode file")
	}

	return fileContent, nil

}
