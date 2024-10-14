package util

import (
	"fmt"
	"strconv"
)

func Input(valueType string) int {
	var userInput string

	fmt.Printf("Enter %s score:\n", valueType)
	fmt.Scan(&userInput)
	value, err := strconv.ParseFloat(userInput, 32)

	if err != nil {
		return 0
	}

	if value > 10.0 {
		return 10
	}

	return int(value)
}

func GetParameters() [4]int {
	art := Input("Art")
	plot := Input("Plot")
	characters := Input("Characters")
	bias := Input("Bias")

	return [4]int{art, plot, characters, bias}
}
