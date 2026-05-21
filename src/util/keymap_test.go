package util

import (
	"testing"

	"github.com/charmbracelet/bubbles/key"
	"github.com/stretchr/testify/assert"
)

func TestShortHelper(t *testing.T) {
	km := AppKeys

	bindings := km.ShortHelp()
	assert.Equal(t, len(bindings), 2)

	tests := []struct {
		name string
		key  key.Binding
		desc string
	}{
		{"help key", km.Help, "toggle help"},
		{"content key", km.Content, "switch content type"},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, bindings[i], tt.key)
			assert.Equal(t, bindings[i].Help().Desc, tt.desc)
		})
	}
}

func TestLongHelper(t *testing.T) {
	km := AppKeys
	bindings := km.FullHelp()

	if !assert.Len(t, bindings, 4) {
		return
	}

	tests := []struct {
		name  string
		key   key.Binding
		count int
		desc  string
	}{
		{"row 1", km.Help, 3, "toggle help"},
		{"row 2", km.Quit, 3, "quit"},
		{"row 3", km.StartOver, 3, "start over"},
		{"row 4", km.Export, 2, "export ratings in json format"},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !assert.Less(t, i, len(bindings), "index out of bounds") {
				return
			}

			got := bindings[i]

			assert.Len(t, got, tt.count)

			if assert.NotEmpty(t, got, "row %d is empty", i) {
				assert.Equal(t, got[0], tt.key)
				assert.Equal(t, got[0].Help().Desc, tt.desc)
			}
		})
	}
}
