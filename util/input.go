package util

import (
	"fmt"
	"strconv"
)

func Input(valueType string) float32 {
	var userInput string

	fmt.Printf("Enter %s score:\n", valueType)
	fmt.Scan(&userInput)
	value, err := strconv.ParseFloat(userInput, 32)

	if err != nil {
		return 0
	}

	return float32(value)
}
