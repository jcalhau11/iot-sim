package device

import (
	"encoding/json"
	"errors"
)

// maybe let front-end decide type
func CustomReflection(target interface{}) (interface{}, string, error) {

	bytesTarget, bytesError := json.Marshal(target)
	if bytesError != nil {
		return nil, "", bytesError
	}

	var options VariedOptions
	unmarchError := json.Unmarshal(bytesTarget, &options)
	if unmarchError == nil {
		return options, "VariedOptions", nil
	}

	var ranges VariedRange
	unmarchRangeError := json.Unmarshal(bytesTarget, ranges)
	if unmarchRangeError == nil {
		return ranges, "VariedRange", nil
	}

	return nil, "", errors.New("Could not reflect this type into Varied options")
}
