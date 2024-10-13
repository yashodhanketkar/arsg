package main

import (
	"fmt"
	"github.com/yashodhanketkar/arsg/util"
)

func Assembly() [2]int {
	score, _ := util.Calculator(1, 1, 1, 1)
	fmt.Printf("Calculator %f\n", score)
	fmt.Printf("Input %d\n", util.Input())
	return [2]int{0, 0}
}

func main() {
	Assembly()
}
