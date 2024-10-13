package main

import (
	"fmt"
	"github.com/yashodhanketkar/arsg/util"
)

func Assembly() [2]int {
	fmt.Printf("Calculator %d\n", util.Calculator())
	fmt.Printf("Input %d\n", util.Input())
	return [2]int{0, 0}
}

func main() {
	Assembly()
}
