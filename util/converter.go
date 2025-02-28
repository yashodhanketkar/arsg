package util

import (
	"math"
)

func SystemCalculator(systemType string, score float32) float32 {
	var converted_score float32

	switch systemType {
	case "Decimal":
		if score < 0.1 {
			return 0.1
		}
		converted_score = score
	case "Integer":
		if score < 1 {
			return 1
		}
		converted_score = float32(int(score))
	case "FivePoint":
		if score < 1 {
			return 1
		}
		converted_score = float32(math.Ceil(float64(score) * 0.5))
	case "Percentage":
		if score < 1 {
			return 1
		}
		converted_score = float32(int(score * 10))
	}

	return converted_score
}
