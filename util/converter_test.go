package util

import (
	"testing"
)

func TestConverters(t *testing.T) {

	t.Run("should return correct values", func(t *testing.T) {

		converterMinTest := []struct {
			score      float32
			systemType string
			output     float32
		}{
			{7.5, "DecimalSystem", 7.5},
			{5.6, "IntegerSystem", 5.0},
			{8.0, "FivePointSystem", 4.0},
			{9, "FivePointSystem", 5.0},
			{3.0, "FivePointSystem", 2.0},
			{2.4, "PercentageSystem", 24},
			{5.4, "PercentageSystem", 54},
		}

		for _, tt := range converterMinTest {

			got := SystemCalculator(tt.systemType, tt.score)
			want := tt.output

			if got != want {
				t.Errorf("got %f, want %f", got, want)
			}
		}
	})

	t.Run("should return highest possible values", func(t *testing.T) {

		converterMinTest := []struct {
			score      float32
			systemType string
			output     float32
		}{
			{9.4, "DecimalSystem", 9.4},
			{9.4, "IntegerSystem", 9},
			{9.4, "FivePointSystem", 5},
			{9.4, "PercentageSystem", 94},
		}

		for _, tt := range converterMinTest {

			got := SystemCalculator(tt.systemType, tt.score)
			want := tt.output

			if got != want {
				t.Errorf("got %f, want %f", got, want)
			}
		}
	})

	t.Run("should return lowest possible values", func(t *testing.T) {

		converterMinTest := []struct {
			score      float32
			systemType string
			output     float32
		}{
			{0, "DecimalSystem", 0.1},
			{0, "IntegerSystem", 1},
			{0, "FivePointSystem", 1},
			{0, "PercentageSystem", 1},
		}

		for _, tt := range converterMinTest {

			got := SystemCalculator(tt.systemType, tt.score)
			want := tt.output

			if got != want {
				t.Errorf("got %f, want %f", got, want)
			}
		}
	})
}
