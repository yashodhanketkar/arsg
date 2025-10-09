package util

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelperHandler(t *testing.T) {
	t.Run("test get parameters", func(t *testing.T) {
		tests := []struct {
			name    string
			handler int
			args    []string
			want    any
		}{
			{
				name:    "Get nil",
				handler: 0,
				args:    []string{},
				want:    nil,
			},
			{
				name:    "Get params default",
				handler: 1,
				args:    []string{},
				want:    []string{"Art/Animation", "Character/Cast", "Plot", "Bias"},
			},
			{
				name:    "Get params via args",
				handler: 1,
				args:    []string{"studio", "genre"},
				want:    []string{"studio", "genre"},
			},
		}

		for _, tt := range tests {
			got := HelperHandler(tt.handler, tt.args...)
			assert.Equal(t, tt.want, got)
		}
	})

	t.Run("test capitalizeFirstLetter", func(t *testing.T) {
		tests := []struct {
			name    string
			handler int
			args    []string
			want    any
		}{
			{
				name:    "no args provided",
				handler: 2,
				args:    []string{},
				want:    fmt.Errorf("No arguments provided"),
			},
			{
				name:    "empty string",
				handler: 2,
				args:    []string{""},
				want:    "",
			},
			{
				name:    "single character string",
				handler: 2,
				args:    []string{"a"},
				want:    "A",
			},
			{
				name:    "mulit-character string",
				handler: 2,
				args:    []string{"abcd"},
				want:    "Abcd",
			},
		}

		for _, tt := range tests {
			got := HelperHandler(tt.handler, tt.args...)
			assert.Equal(t, tt.want, got)
		}
	})
}

func TestCalculator(t *testing.T) {

	t.Run("should adjust and round the score", func(t *testing.T) {

		adjusterTests := []struct {
			target float32
			want   float32
		}{
			{10.0, 9.4},
			{6.7, 6.3},
			{0, 0.0},
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
		_, err := Calculator([]float32{0, 0, 0, 0}...)

		if err == nil {
			t.Error("should throw error")
		}
	})

	t.Run("should return correct scores", func(t *testing.T) {

		calculatorTests := []struct {
			parameters []float32
			want       float32
		}{
			{[]float32{5, 6, 3, 1}, 4.1},
			{[]float32{10, 10, 10, 10}, 9.4},
		}

		for _, tt := range calculatorTests {
			got, _ := Calculator(tt.parameters...)

			if got != tt.want {
				t.Errorf("want %f, got %f", tt.want, got)
			}
		}
	})
}

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
		want := float32(tt.expected)
		got := FloatParser(tt.value)
		if got != want {
			t.Errorf("want %v; got %v", want, got)
		}
	}
}

func TestInput(t *testing.T) {

	t.Run("should return 0", func(t *testing.T) {
		want := float32(0)
		originalStdin := os.Stdin

		defer func() { os.Stdin = originalStdin }()
		r, w, _ := os.Pipe()
		w.Write([]byte("\n"))
		w.Close()
		os.Stdin = r

		got := Input("Plot")
		if got != want {
			t.Errorf("want %v; got %v", want, got)
		}
	})

	t.Run("should return 10", func(t *testing.T) {
		want := float32(10.0)
		originalStdin := os.Stdin

		defer func() { os.Stdin = originalStdin }()
		r, w, _ := os.Pipe()
		w.Write([]byte("30\n"))
		w.Close()
		os.Stdin = r

		got := Input("Plot")
		if got != want {
			t.Errorf("want %v; got %v", want, got)
		}
	})

	t.Run("should return correct result", func(t *testing.T) {
		want := float32(3.45)
		originalStdin := os.Stdin

		defer func() { os.Stdin = originalStdin }()
		r, w, _ := os.Pipe()
		w.Write([]byte("3.45\n"))
		w.Close()
		os.Stdin = r

		got := Input("Plot")
		if got != want {
			t.Errorf("want %v; got %v", want, got)
		}
	})
}

func TestGetParameters(t *testing.T) {
	t.Run("should return [0, 0, 0, 0]", func(t *testing.T) {
		got := GetParameters()
		want := []float32{0, 0, 0, 0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("should return correct result", func(t *testing.T) {
		want := []float32{2.3, 4.2, 1.6, 9.2}
		originalStdin := os.Stdin

		defer func() { os.Stdin = originalStdin }()

		r, w, _ := os.Pipe()
		w.Write([]byte("2.3\n"))
		w.Write([]byte("4.2\n"))
		w.Write([]byte("1.6\n"))
		w.Write([]byte("9.2\n"))
		w.Close()

		os.Stdin = r

		got := GetParameters()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v; got %v", want, got)
		}
	})
}

func TestGetNumericInput(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{"123.45", "123.45"},
		{"a", ""},
		{"1a23", "123"},
		{"a123", "123"},
	}

	for i, tt := range testCases {
		got := GetNumericInput(tt.input)
		want := tt.output

		if got != want {
			t.Errorf("[%d] want %v; got %v", i+1, want, got)
		}
	}
}
