package main

import (
	"encoding/json"
	"fmt"
	"github.com/jcalhau11/iot-sim/oven"
)

func main() {

	file := &oven.File{Path: "phxs/testy.json"}

	comp, compErro := file.ReloadDevice()

	fmt.Println(compErro)

	fmt.Println(comp)

	comp.Telemetry()

	t, _ := json.Marshal(comp.Static)

	fmt.Println(string(t))

}
