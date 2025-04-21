package util

import (
	"fmt"
	"strconv"
	"strings"
)

func Input(valueType string) float32 {
	var userInput string

	fmt.Printf("Enter %s score:\n", valueType)
	fmt.Scan(&userInput)
	value, err := strconv.ParseFloat(userInput, 32)

	if err != nil {
		return 0
	}

	if value > 10.0 {
		return float32(10.0)
	} else {
		return float32(value)
	}
}

func GetParameters() [4]float32 {
	art := Input("Art")
	plot := Input("Plot")
	characters := Input("Characters")
	bias := Input("Bias")

	return [4]float32{art, plot, characters, bias}
}

func GetNumericInput(str string) string {
	var inputBuilder strings.Builder

	for _, r := range str {
		if strings.ContainsRune("0123456789.", r) {
			inputBuilder.WriteRune(r)
		}
	}

	return inputBuilder.String()
}
