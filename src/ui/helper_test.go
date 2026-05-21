package ui

import (
	"fmt"
	"testing"

	"github.com/atotto/clipboard"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/assert"
	"github.com/yashodhanketkar/arsg/src/db"
)

func TestSetupParameters(t *testing.T) {
	tests := []struct {
		name           string
		params         []string
		expcetedParams []string
		expcetedErr    error
	}{
		{
			"test default parameters",
			[]string{},
			[]string{"Art/Animation", "Character/Cast", "Plot", "Bias"},
			nil,
		}, {
			"test custom parameters",
			[]string{"test", "test2"},
			[]string{"Test", "Test2"},
			nil,
		}, {
			"test invalid parameters - single - empty string",
			[]string{""},
			nil,
			fmt.Errorf("Empty string provided"),
		}, {
			"test invalid parameters - second - empty string",
			[]string{"test", ""},
			nil,
			fmt.Errorf("Empty string provided"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := setupParameters(tt.params...)
			assert.Equal(t, tt.expcetedParams, got)
			assert.Equal(t, tt.expcetedErr, err)
		})
	}
}

func TestSetFoucs(t *testing.T) {
	m := helperInitilizeModel(t)
	assert.NotEqual(t, focusedStyle, m.form.inputs[1].PromptStyle)
	assert.Equal(t, m.form.focusIndex, 0)
	m.setFocus(1)
	assert.Equal(t, focusedStyle, m.form.inputs[1].PromptStyle)
}

func TestIsNumeric(t *testing.T) {
	m := helperInitilizeModel(t)
	m.form.focusIndex = 6
	assert.False(t, m.isNumeric())
	m.form.focusIndex = 1
	assert.True(t, m.isNumeric())
}

func TestResetInputs(t *testing.T) {
	m := helperInitilizeModel(t)
	mockValue(t, m)
	m.calculateScore() // set m.state.score to 6.8
	m.resetInputs()
	assert.Equal(t, 0, m.form.focusIndex)
	assert.Equal(t, float32(0.0), m.state.score)
	for i := range 5 {
		assert.Equal(t, "", m.form.inputs[i].Value())
	}
	assert.Equal(t, float32(0.0), m.state.score)
}

func TestCalculateScore(t *testing.T) {
	m := helperInitilizeModel(t)
	mockValue(t, m)

	// check correct score
	m.calculateScore()
	assert.Equal(t, float32(7.2), m.state.score)

	// check max allowed score
	for _, v := range []string{"10", "15"} {
		m.form.inputs[4].SetValue(v)
		m.calculateScore()
		assert.Equal(t, float32(7.7), m.state.score)
	}

	// check all 0 score
	m.form.inputs[1].SetValue("0")
	m.form.inputs[2].SetValue("0")
	m.form.inputs[3].SetValue("0")
	m.form.inputs[4].SetValue("0")
	m.calculateScore()
	assert.Equal(t, float32(0.0), m.state.score)
}

func TestScoreToClipboard(t *testing.T) {
	m := helperInitilizeModel(t)

	mockValue(t, m)

	m.calculateScore()
	m.copyToClipboard()

	v, _ := clipboard.ReadAll()
	assert.Equal(t, "7.2", v)

	clipboard.WriteAll("Test Name")
	m.pasteFromClipbaord()
	assert.Equal(t, "Test Name", m.form.inputs[0].Value())
}

func TestUpdateInputs(t *testing.T) {
	m := helperInitilizeModel(t)
	m.updateInputs(tea.KeyMsg{Type: tea.KeyEnter})
}

func TestPrepareRatings(t *testing.T) {
	m := helperInitilizeModel(t)
	mockValue(t, m)
	m.calculateScore()
	got := m.prepareRating()

	expected := db.Rating{
		Name:     "test",
		Art:      8.5,
		Support:  7.5,
		Plot:     6.5,
		Bias:     5.5,
		Rating:   "7.2",
		Comments: "comment",
	}

	assert.Equal(t, expected, got)
}

func TestSwitchContent(t *testing.T) {
	m := helperInitilizeModel(t)

	assert.Equal(t, "anime", m.viewer.contentType)

	for _, option := range []string{"manga", "lightnovel", "anime"} {
		assert.Equal(t, option, m.toggleContentType())
	}
}
