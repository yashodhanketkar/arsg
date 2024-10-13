package main

import (
	"testing"
)

func TestAssembly(t *testing.T) {
	got := Assembly()
	want := [2]int{0, 0}

	if got != want {
		t.Errorf("want %d, got %d", want, got)

	}
}

func TestMain(t *testing.T) {
	main()
}
