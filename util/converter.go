package util

import (
	"math"
	"strconv"
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

func FloatParser(value string) float32 {
	floatValue, error := strconv.ParseFloat(value, 64)
	if error != nil {
		return 0
	}
	parsedFloat := float32(floatValue)

	if parsedFloat < 0.0 {
		parsedFloat = 0.0
	} else if parsedFloat > 10.0 {
		parsedFloat = 10.0
	}

	return parsedFloat
}
