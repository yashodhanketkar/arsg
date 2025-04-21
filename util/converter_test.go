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
			{7.5, "Decimal", 7.5},
			{5.6, "Integer", 5.0},
			{8.0, "FivePoint", 4.0},
			{9, "FivePoint", 5.0},
			{3.0, "FivePoint", 2.0},
			{2.4, "Percentage", 24},
			{5.4, "Percentage", 54},
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
			{9.4, "Decimal", 9.4},
			{9.4, "Integer", 9},
			{9.4, "FivePoint", 5},
			{9.4, "Percentage", 94},
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
			{0, "Decimal", 0.1},
			{0, "Integer", 1},
			{0, "FivePoint", 1},
			{0, "Percentage", 1},
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

func TestFloatParser(t *testing.T) {
	t.Run("should return 0", func(t *testing.T) {
		want := float32(0)
		got := FloatParser("a")
		if got != want {
			t.Errorf("want %v; got %v", want, got)
		}
	})

	t.Run("should return 10", func(t *testing.T) {
		want := float32(10.0)
		got := FloatParser("10")
		if got != want {
			t.Errorf("want %v; got %v", want, got)
		}
	})

	t.Run("should return correct result", func(t *testing.T) {
		want := float32(3.45)
		got := FloatParser("3.45")
		if got != want {
			t.Errorf("want %v; got %v", want, got)
		}
	})
}
