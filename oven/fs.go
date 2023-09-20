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

func (file *File) WriteFile(content interface{}) error {

	document, documentErr := os.OpenFile(file.Path, os.O_WRONLY|os.O_CREATE, 0644)
	if documentErr != nil {
		return documentErr
	}

	defer document.Close()

	if err := document.Truncate(0); err != nil {
		return err
	}
	if _, err := document.Seek(0, 0); err != nil {
		return err
	}

	bytesContent, bytesContentError := json.Marshal(content)

	if bytesContentError != nil {
		return bytesContentError
	}

	if _, writeError := document.WriteString(string(bytesContent)); writeError != nil {
		return writeError
	}

	return nil
}
