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

func TestCalculator(t *testing.T) {
	config := mockConfig(t)

	t.Run("should throw error", func(t *testing.T) {
		got, err := Calculator(&config, []float32{0, 0, 0, 0}...)

		assert.Error(t, err)
		assert.Equal(t, float32(0.0), got)
	})

	t.Run("should return correct scores", func(t *testing.T) {
		calculatorTests := []struct {
			parameters []float32
			want       float32
		}{
			{[]float32{1, 2, 3, 4}, 2.3},
			{[]float32{5, 4, 3, 1}, 3.7},
			{[]float32{5, 5.5, 5, 5}, 5.2}, {[]float32{10, 10, 10, 10}, 10.0},
		}

		for _, tt := range calculatorTests {
			got, err := Calculator(&config, tt.parameters...)

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
			got, err := Calculator(&config, tt.parameters...)

			assert.Error(t, err)
			assert.Equal(t, err.Error(), tt.err)
			assert.Equal(t, float32(0.0), got)
		}
	})

}

func TestRounder(t *testing.T) {
	t.Run("should adjust and round the score", func(t *testing.T) {

		adjusterTests := []struct {
			target float32
			want   float32
		}{
			{10.0, 10.0},
			{6.7, 6.7},
			{0, 0.0},
		}

		for _, tt := range adjusterTests {
			got := rounder(tt.target)
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
				got := rounder(tt.target)
				assert.Equal(t, tt.want, got)
			})
		}
	})
}
