package main

import (
	"fmt"
	"github.com/yashodhanketkar/arsg/util"
)

func Assembly() [2]int {
	parameters := util.GetParameters()
	score, _ := util.Calculator(parameters)
	fmt.Printf("Calculator %f\n", score)
	return [2]int{0, 0}
}

func main() {
	Assembly()
}
