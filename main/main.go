package main

import (
	"fmt"

	"github.com/yashodhanketkar/arsg/util"
)

func mainLoop() []float32 {
	choice := "Y"
	var res []float32

	for choice == "Y" || choice == "y" {
		choice = "N"
		parameters := util.GetParameters()

		if score, err := util.Calculator(parameters); err == nil {
			res = append(res, score)
			fmt.Printf("Generated rating for this item is %v\n", score)
		}
		fmt.Println("To continue enter y/Y")
		fmt.Scan(&choice)
	}

	return res
}

func main() {
	mainLoop()
}
