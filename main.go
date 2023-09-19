package main

import (
	"fmt"
	"github.com/jcalhau11/iot-sim/oven"
)

func main() {

	file := &oven.File{Path: "phxs/testy.json"}

	fmt.Println(file.ReadFile())
}
