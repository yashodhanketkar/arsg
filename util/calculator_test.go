package util

import "testing"

func TestCalculator(t *testing.T) {

	t.Run("should adjust and round the score", func(t *testing.T) {

		adjusterTests := []struct {
			target float32
			want   float32
		}{
			{10.0, 9.4},
			{6.7, 6.3},
		}

		for _, tt := range adjusterTests {
			got := adjuster(tt.target)

			if got != tt.want {
				t.Errorf("want %f, got %f", tt.want, got)
			}
		}
	})

	t.Run("should round by 1 digit", func(t *testing.T) {

		rounderTests := []struct {
			target float32
			want   float32
		}{
			{12.363, 12.4},
			{12.343, 12.3},
		}

		for _, tt := range rounderTests {
			got := rounder(tt.target)

			if got != tt.want {
				t.Errorf("want %f, got %f", tt.want, got)
			}
		}
	})

	t.Run("should throw error", func(t *testing.T) {
		_, err := Calculator(0, 0, 0, 0)

		if err == nil {
			t.Error("should throw error")
		}
	})

	t.Run("should return correct scores", func(t *testing.T) {

		calculatorTests := []struct {
			art        int
			plot       int
			characters int
			bias       int
			want       float32
		}{
			{5, 6, 3, 1, 4.1},
			{10, 10, 10, 10, 9.4},
		}

		for _, tt := range calculatorTests {
			got, _ := Calculator(tt.art, tt.plot, tt.characters, tt.bias)

			if got != tt.want {
				t.Errorf("want %f, got %f", tt.want, got)
			}
		}
	})
}