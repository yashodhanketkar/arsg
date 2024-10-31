package main

import (
	"os"
	"reflect"
	"testing"
)

func TestMainLoop(t *testing.T) {

	t.Run("should return correct result", func(t *testing.T) {
		want := []float32{3.2}
		originalStdin := os.Stdin

		defer func() { os.Stdin = originalStdin }()

		r, w, _ := os.Pipe()
		w.Write([]byte("2.3\n"))
		w.Write([]byte("4.2\n"))
		w.Write([]byte("1.6\n"))
		w.Write([]byte("9.2\n"))
		w.Write([]byte("N\n"))
		w.Close()
		os.Stdin = r

		got := mainLoop()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %f, got %f", want, got)
		}
	})

	t.Run("should return 9.4", func(t *testing.T) {
		want := []float32{9.4}
		originalStdin := os.Stdin

		defer func() { os.Stdin = originalStdin }()

		r, w, _ := os.Pipe()
		w.Write([]byte("13\n"))
		w.Write([]byte("10\n"))
		w.Write([]byte("16\n"))
		w.Write([]byte("92\n"))
		w.Write([]byte("N\n"))
		w.Close()
		os.Stdin = r

		got := mainLoop()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %f, got %f", want, got)
		}
	})

	t.Run("should return 0.1", func(t *testing.T) {
		want := []float32{0.1}
		originalStdin := os.Stdin

		defer func() { os.Stdin = originalStdin }()

		r, w, _ := os.Pipe()
		w.Write([]byte("0.1\n"))
		w.Write([]byte("0\n"))
		w.Write([]byte("0\n"))
		w.Write([]byte("0\n"))
		w.Write([]byte("N\n"))
		w.Close()
		os.Stdin = r

		got := mainLoop()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %f, got %f", want, got)
		}
	})

	t.Run("should run twice and return 0.1 both times", func(t *testing.T) {
		want := []float32{0.1, 0.1}
		originalStdin := os.Stdin

		defer func() { os.Stdin = originalStdin }()

		r, w, _ := os.Pipe()
		w.Write([]byte("0.1\n"))
		w.Write([]byte("0\n"))
		w.Write([]byte("0\n"))
		w.Write([]byte("0\n"))
		w.Write([]byte("Y\n"))
		w.Write([]byte("0.1\n"))
		w.Write([]byte("0\n"))
		w.Write([]byte("0\n"))
		w.Write([]byte("0\n"))
		w.Write([]byte("N\n"))
		w.Close()
		os.Stdin = r

		got := mainLoop()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %f, got %f", want, got)
		}
	})
}
