package oven

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jcalhau11/iot-sim/device"
	"os"
)

type File struct {
	Path string
}

func (file *File) ReadFile() (map[string]interface{}, error) {
	fl, flErr := os.Open(file.Path)
	if flErr != nil {
		return nil, errors.New("File not found")
	}

	defer fl.Close()

	var fileContent map[string]interface{}
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

func (file *File) ReloadDevice() (*device.Device, error) {
	fileContent, fileContentError := file.ReadFile()
	if fileContentError != nil {
		return nil, fileContentError
	}

	deviceVarieds := make([]device.Varied, 0)

	varied := fileContent["varies"]

	for _, value := range varied.([]interface{}) {
		bytes, errBytes := json.Marshal(value)
		if errBytes != nil {
			return nil, errBytes
		}

		var def device.DefaultVaried

		_ = json.Unmarshal(bytes, &def)

		varies, errVaries := selectOnType(def.Type, value)
		if errVaries != nil {
			return nil, errVaries
		}

		deviceVarieds = append(deviceVarieds, varies)
	}

	return &device.Device{
		Name:   fileContent["name"].(string),
		Type:   fileContent["type"].(string),
		Static: fileContent["static"].(map[string]interface{}),
		Varies: deviceVarieds,
	}, nil

}

func selectOnType(typ string, val interface{}) (device.Varied, error) {
	if typ == "variedOptions" {
		byteArray, _ := json.Marshal(val)
		var variedOption device.VariedOptions
		if err := json.Unmarshal(byteArray, &variedOption); err != nil {
			return nil, err
		}

		return variedOption, nil
	}

	if typ == "variedRange" {
		byteArray, _ := json.Marshal(val)
		var variedRange device.VariedRange
		if err := json.Unmarshal(byteArray, &variedRange); err != nil {
			return nil, err
		}

		return variedRange, nil
	}

	return nil, errors.New("Invalid type")
}
