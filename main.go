package main

import (
	"encoding/json"
	"fmt"
	dev "github.com/jcalhau11/iot-sim/device"
	"github.com/jcalhau11/iot-sim/oven"
)

func main() {

	file := &oven.File{Path: "phxs/testy.json"}

	fileInterface, _ := file.ReadFile()

	jsonSt, _ := json.Marshal(fileInterface)

	var device dev.Device

	err := json.Unmarshal(jsonSt, &device)

	fmt.Println(err)

	var opts []interface{}

	opts = append(opts, "on")
	opts = append(opts, "off")

	device.ForceOptions(dev.VariedOptions{
		Path:  "power.is_connected",
		Opts:  opts,
		Probs: []string{"25", "75"},
	})

	fmt.Println(device)

}
