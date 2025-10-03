package ui

import "fmt"

type item struct {
	id         int
	title      string
	desc       string
	parameters [4]float32
	score      string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return fmt.Sprintf("%s - %+v", i.score, i.parameters) }
func (i item) FilterValue() string { return i.title }
