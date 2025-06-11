package ui

import (
	"testing"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/assert"
	"github.com/yashodhanketkar/arsg/db"
)

func TestInit(t *testing.T) {
	m := initialModel()
	assert.IsType(t, model{}, m)
	cmd := m.Init()
	_, ok := cmd().(tea.BatchMsg)
	assert.True(t, ok)

	assert.Equal(t, m.scoreMode, 0)
	assert.Equal(t, m.help, help.New())
	assert.Equal(t, m.view, 0)

	for i, v := range []string{
		"Name",
		"Art/Animation",
		"Character/Cast",
		"Plot",
		"Bias",
		"Comments",
	} {
		assert.Equal(t, v, m.inputs[i].Placeholder)
	}
}

func TestHelpers(t *testing.T) {
	t.Run("test setFocus", func(t *testing.T) {
		m := initialModel()
		assert.NotEqual(t, focusedStyle, m.inputs[1].PromptStyle)
		assert.Equal(t, m.focusIndex, 0)
		m.setFocus(1)
		assert.Equal(t, focusedStyle, m.inputs[1].PromptStyle)
	})

	t.Run("test isNumeric", func(t *testing.T) {
		m := initialModel()
		m.focusIndex = 6
		assert.False(t, m.isNumeric())
		m.focusIndex = 1
		assert.True(t, m.isNumeric())
	})

	t.Run("test resetInputs", func(t *testing.T) {
		m := initialModel()
		mockValue(t, m)
		m.calculateScore() // set m.score to 6.8
		m.resetInputs()
		assert.Equal(t, 0, m.focusIndex)
		assert.Equal(t, float32(0.0), m.score)
		for i := range 5 {
			assert.Equal(t, "", m.inputs[i].Value())
		}
		assert.Equal(t, float32(0.0), m.score)
	})

	t.Run("test calculateScore", func(t *testing.T) {
		m := initialModel()
		mockValue(t, m)

		// check correct score
		m.calculateScore()
		assert.Equal(t, float32(6.8), m.score)

		// check max allowed score
		for _, v := range []string{"10", "15"} {
			m.inputs[4].SetValue(v)
			m.calculateScore()
			assert.Equal(t, float32(7.2), m.score)
		}

		// check all 0 score
		m.inputs[1].SetValue("0")
		m.inputs[2].SetValue("0")
		m.inputs[3].SetValue("0")
		m.inputs[4].SetValue("0")
		m.calculateScore()
		assert.Equal(t, float32(0.0), m.score)
	})

	t.Run("test scoreToClipboard", func(t *testing.T) {
		m := initialModel()
		mockValue(t, m)
		m.calculateScore()
		m.copyToClipboard()
		v, _ := clipboard.ReadAll()
		assert.Equal(t, "6.8", v)
		clipboard.WriteAll("Test Name")
		m.pasteFromClipbaord()
		assert.Equal(t, "Test Name", m.inputs[0].Value())
	})

	t.Run("test updateInputs", func(t *testing.T) {
		m := initialModel()
		m.updateInputs(tea.KeyMsg{Type: tea.KeyEnter})
	})

	t.Run("test prepare ratings", func(t *testing.T) {
		m := initialModel()
		mockValue(t, m)
		m.calculateScore()
		got := m.prepareRating()

		expected := db.Rating{
			Name:     "test",
			Art:      8.5,
			Support:  7.5,
			Plot:     6.5,
			Bias:     5.5,
			Rating:   "6.8",
			Comments: "comment",
		}

		assert.Equal(t, expected, got)
	})
}

func TestButtonCommands(t *testing.T) {

	t.Run("test default behaviour", func(t *testing.T) {
		m := initialModel()
		nm, cmd := m.buttonCommands()
		assert.Nil(t, cmd)

		switch nmt := nm.(type) {
		case model:
			assert.Equal(t, m.focusIndex+1, nmt.focusIndex)
		default:
			t.Errorf("incorrect model recieved. %T", nm)
		}
	})

	t.Run("test reset button", func(t *testing.T) {
		m := initialModel()
		mockValue(t, m)
		m.calculateScore()
		m.focusIndex = len(m.inputs) + 1
		nm, cmd := m.buttonCommands()
		assert.Nil(t, cmd)

		switch nm := nm.(type) {
		case model:
			assert.Equal(t, 0, nm.focusIndex)
			assert.Equal(t, float32(0.0), nm.score)
		default:
			t.Errorf("incorrect model recieved. %T", nm)
		}

		t.Run("test quit button", func(t *testing.T) {
			m := initialModel()
			m.focusIndex = len(m.inputs) + 2
			_, cmd := m.buttonCommands()
			_, ok := cmd().(tea.QuitMsg)
			assert.True(t, ok)
		})
	})
}

func mockValue(t *testing.T, m model) {
	t.Helper()

	m.inputs[0].SetValue("test")
	m.inputs[1].SetValue("8.5")
	m.inputs[2].SetValue("7.5")
	m.inputs[3].SetValue("6.5")
	m.inputs[4].SetValue("5.5")
	m.inputs[5].SetValue("comment")
}
