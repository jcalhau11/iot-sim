package device

import (
	"errors"
	"strings"
)

func (device *Device) ChangeAttr(path string, value any) error {
	paths := strings.Split(path, ".")
	var nestedList []map[string]interface{}
	static := device.Static
	nestedList = append(nestedList, static)

	for _, key := range paths {
		val, exists := static[key]
		if !exists {
			return errors.New("Invalid path")
		}

		if nested, isNested := val.(map[string]interface{}); isNested {
			nestedList = append(nestedList, nested)
		} else {
			nestedList[len(nestedList)-1][key] = value
		}
	}

	if len(nestedList) > 1 {
		for i := len(nestedList) - 1; i > 0; i-- {
			nestedList[i-1][paths[i-1]] = nestedList[i][paths[i]]
		}
	}

	device.Static = nestedList[0]

	return nil
}
