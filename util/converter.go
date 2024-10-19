package util

import (
	"fmt"
	"math"
)

func SystemCalculator(systemType string, score float32) float32 {
	var converted_score float32

	switch systemType {
	case "DecimalSystem":
		if score < 0.1 {
			return 0.1
		}
		converted_score = score
	case "IntegerSystem":
		if score < 0.1 {
			return 1
		}
		converted_score = float32(int(score))
	case "FivePointSystem":
		if score < 0.1 {
			return 1
		}
		converted_score = float32(math.Ceil(float64(score) * 0.5))
	case "PercentageSystem":
		if score < 0.1 {
			return 1
		}
		converted_score = float32(int(score * 10))
	}

	fmt.Printf("Debug 1: Score is %f\n", converted_score)

	return converted_score
}
