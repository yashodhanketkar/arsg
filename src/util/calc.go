package util

import (
	"fmt"
	"math"
)

func rounder(score float32, limiter float32) float32 {
	return float32(math.Round(float64(score*limiter)) / 10)
}

func Calculator(config *ConfigType, limiter float32, args ...float32) (float32, error) {
	if len(args) == 0 {
		return 0, fmt.Errorf("No input values provided")
	}

	if limiter > 10.0 {
		limiter = 10
	} else if limiter < 1.0 {
		limiter = 1
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

	return rounder(weightedSum/maxTotal, limiter), nil
}
