package util

import (
	"os"
	"testing"
)

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

	t.Run("should return correct result", func(t *testing.T) {
		want := float32(123.45)
		originalStdin := os.Stdin

		defer func() { os.Stdin = originalStdin }()
		r, w, _ := os.Pipe()
		w.Write([]byte("123.45\n"))
		w.Close()
		os.Stdin = r

		got := Input("Plot")
		if got != want {
			t.Errorf("want %v; got %v", want, got)
		}
	})
}
