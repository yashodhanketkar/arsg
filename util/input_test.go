package util

import (
	"os"
	"testing"
)

func TestInput(t *testing.T) {

	t.Run("should return 0", func(t *testing.T) {
		want := 0
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
		want := 10
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
		want := 3
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
		want := [4]int{0, 0, 0, 0}

		if got != want {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("should return correct result", func(t *testing.T) {
		want := [4]int{2, 4, 1, 9}
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

		if got != want {
			t.Errorf("want %v; got %v", want, got)
		}
	})
}
