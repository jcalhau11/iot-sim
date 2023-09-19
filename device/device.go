package device

import (
	"errors"
	"strings"
)

func (device *Device) GetAttribute(path string) (any, error) {

	paths := strings.Split(path, ".")
	baseLine := device.Static

	for _, key := range paths {
		next, exists := baseLine[key]
		if !exists {
			return nil, errors.New("Invalid path, path does not exist")
		}

		if nested, ok := next.(map[string]interface{}); ok {
			baseLine = nested
		} else {
			return next, nil
		}
	}

	return nil, errors.New("Invalid path, path does not exist")
}

func (device *Device) UpdateAttribute(path string, value any) error {
	return device.ChangeAttr(path, value)
}
