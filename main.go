package main

import (
	"encoding/json"
	"fmt"
	device "github.com/jcalhau11/iot-sim/device"
	"github.com/jcalhau11/iot-sim/oven"
)

func main() {

	file := &oven.File{Path: "phxs/testy.json"}

	fileInterface, _ := file.ReadFile()

	jsonSt, _ := json.Marshal(fileInterface)

	var device device.Device

	err := json.Unmarshal(jsonSt, &device)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(device)

	writeError := device.ChangeAttr("humidity.c.d", "changed")

	errF := file.WriteFile(device)

	fmt.Println(errF)

	fmt.Println(writeError)

	fmt.Println(device)
}
