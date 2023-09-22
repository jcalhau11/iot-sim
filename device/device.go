package device

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
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

func (variation VariedOptions) Force(device *Device) error {

	optionsProbailityMx := make(map[interface{}]int)
	var selectedKey interface{}
	for idx, opts := range variation.Opts {
		integer, _ := strconv.Atoi(variation.Probs[idx])
		optionsProbailityMx[opts] = integer
	}

	randomSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randomSource)

	selectingProability := random.Intn(101)

	threshold := 0

	for option, optionProbability := range optionsProbailityMx {
		upperLimit := threshold + optionProbability

		if selectingProability >= threshold && selectingProability < upperLimit {
			selectedKey = option
		}
	}

	if selectedKey == nil {
		return errors.New("Error occured options not selected")
	}

	device.ChangeAttr(variation.Path, selectedKey)

	return nil

}

func (variedRange VariedRange) Force(device *Device) error {
	max := variedRange.Range[1]
	min := variedRange.Range[0]

	rand.NewSource(time.Now().UnixNano())
	selecteRange := rand.Intn(max-min) + min

	err := device.ChangeAttr(variedRange.Path, selecteRange)

	return err
}

func (device *Device) Telemetry() error {
	for _, action := range device.Varies {
		action.Force(device)
	}

	return nil

}
