package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockConfig(t *testing.T) ConfigType {
	t.Helper()
	config := ConfigType{
		Parameters: []ParamType{
			{"art": 25},
			{"plot": 35},
			{"character": 30},
			{"bias": 10},
		},
	}

	return config
}

func TestCalculatorLimter(t *testing.T) {
	config := mockConfig(t)

	tests := []struct {
		name    string
		limiter float32
		want    float32
	}{
		{"should return for 10.0 limter", 10.0, 2.3},
		{"should return for 1.0 limter", 1.0, 0.2},
		{"should return for 5.0 limter", 5.0, 1.1},
		{"should return for 0.1 limter", 0.1, 0.2},
		{"should return for 0 limter", 0.0, 0.2},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculator(&config, tt.limiter, []float32{1, 2, 3, 4}...)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCalculator(t *testing.T) {
	config := mockConfig(t)

	t.Run("should throw error", func(t *testing.T) {
		got, err := Calculator(&config, 100.0, []float32{0, 0, 0, 0}...)

		assert.Error(t, err)
		assert.Equal(t, float32(0.0), got)
	})

	t.Run("should return correct scores", func(t *testing.T) {
		calculatorTests := []struct {
			parameters []float32
			limiter    float32
			want       float32
		}{
			{[]float32{1, 2, 3, 4}, 10, 2.3},
			{[]float32{5, 4, 3, 1}, 10, 3.7},
			{[]float32{5, 5.5, 5, 5}, 10, 5.2},
			{[]float32{10, 10, 10, 10}, 10, 10.0},
			{[]float32{10, 10, 10, 10}, 9.4, 9.4},
		}

		for _, tt := range calculatorTests {
			got, err := Calculator(&config, tt.limiter, tt.parameters...)

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		}
	})

	t.Run("should result in error", func(t *testing.T) {
		calculatorTests := []struct {
			parameters []float32
			err        string
		}{
			{[]float32{}, "No input values provided"},
			{[]float32{1}, "Invalid number of inputs provided. 1[len(args)] != 4[len(weights)]"},
		}

		for _, tt := range calculatorTests {
			got, err := Calculator(&config, 10.0, tt.parameters...)

			assert.Error(t, err)
			assert.Equal(t, err.Error(), tt.err)
			assert.Equal(t, float32(0.0), got)
		}
	})

}

func TestRounder(t *testing.T) {
	t.Run("should adjust and round the score", func(t *testing.T) {

		adjusterTests := []struct {
			target  float32
			limiter float32
			want    float32
		}{
			{10.0, 10.0, 10.0},
			{6.7, 10.0, 6.7},
			{0, 10.0, 0.0},
			{10.0, 9.0, 9.0},
			{10.0, 1.0, 1.0},
			{6.7, 1.0, 0.7},
		}

		for _, tt := range adjusterTests {
			got := rounder(tt.target, tt.limiter)
			assert.Equal(t, tt.want, got)
		}
	})

	t.Run("should round by 1 digit", func(t *testing.T) {
		rounderTests := []struct {
			name   string
			target float32
			want   float32
		}{
			{"round up", 12.363, 12.4},
			{"round down", 12.343, 12.3},
		}

		for _, tt := range rounderTests {
			t.Run(tt.name, func(t *testing.T) {
				got := rounder(tt.target, 10.0)
				assert.Equal(t, tt.want, got)
			})
		}
	})
}
