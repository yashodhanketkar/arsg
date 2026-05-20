package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
			assert.Equal(t, tt.output, got)
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
			assert.Equal(t, tt.output, got)
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
			assert.Equal(t, tt.output, got)
		}
	})
}

func TestFloatParser(t *testing.T) {
	tests := []struct {
		value    string
		expected float32
	}{
		{"a", 0.0},
		{"", 0.0},
		{"10", 10.0},
		{"0", 0.0},
		{"-5", 0.0},
		{"-0.001", 0.0},
		{"15", 10.0},
		{"10.001", 10.0},
	}

	for _, tt := range tests {
		got := FloatParser(tt.value)
		assert.Equal(t, float32(tt.expected), got)
	}
}

func TestGetNumericInput(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"123.45", "123.45"},
		{"a", ""},
		{"1a23", "123"},
		{"a123", "123"},
	}

	for _, tt := range tests {
		got := GetNumericInput(tt.input)
		assert.Equal(t, tt.output, got)
	}
}
