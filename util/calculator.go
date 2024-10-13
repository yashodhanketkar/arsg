package util

import (
	"errors"
	"math"
)

func rounder(score float32) float32 {
	return float32(math.Round(float64(score*10)) / 10)
}

func adjuster(score float32) float32 {
	return rounder((score / 10) * 9.4)
}

func Calculator(art, plot, characters, bias int) (float32, error) {

	if art == 0 || plot == 0 || characters == 0 || bias == 0 {
		return 0, errors.New("Invalid input")
	}

	score := float32(art*25+plot*35+characters*30+bias*10) / 100

	return adjuster(score), nil
}
