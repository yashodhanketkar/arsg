package util

import (
	"fmt"
	"math"
)

func rounder(score float32) float32 {
	return float32(math.Round(float64(score*10)) / 10)
}

func Calculator(config *ConfigType, args ...float32) (float32, error) {
	if len(args) == 0 {
		return 0, fmt.Errorf("No input values provided")
	}

	_, weights := GetParams(config)
	if len(weights) != len(args) {
		return 0, fmt.Errorf(
			"Invalid number of inputs provided. %d[len(args)] != %d[len(weights)]",
			len(args), len(weights),
		)
	}

	var (
		weightedSum float32
		maxTotal    float32
		hasValue    bool
	)

	for i, arg := range args {
		if arg != 0 {
			hasValue = true
		}
		weightedSum += (args[i] * float32(weights[i]))
		maxTotal += float32(weights[i])
	}

	if !hasValue {
		return 0, fmt.Errorf("All zero values provided")
	}

	return rounder(weightedSum / maxTotal), nil
}
