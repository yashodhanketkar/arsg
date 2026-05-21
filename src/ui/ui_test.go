package ui

import (
	"database/sql"
	"testing"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/assert"
	"github.com/yashodhanketkar/arsg/src/db"
	"github.com/yashodhanketkar/arsg/src/util"
)

func helperInitilizeModel(t *testing.T, args ...string) model {
	t.Helper()

	memDB := util.MockDB(t)
	origDB := db.ConnectDB
	db.ConnectDB = func(_ string) *sql.DB {
		return memDB
	}
	t.Cleanup(func() {
		db.ConnectDB = origDB
	})

	setupParameters()

	return initialModel(memDB, 10.0)
}

func mockValue(t *testing.T, m model) {
	t.Helper()

	m.form.inputs[0].SetValue("test")
	m.form.inputs[1].SetValue("8.5")
	m.form.inputs[2].SetValue("7.5")
	m.form.inputs[3].SetValue("6.5")
	m.form.inputs[4].SetValue("5.5")
	m.form.inputs[5].SetValue("comment")
}

func TestInit(t *testing.T) {
	m := helperInitilizeModel(t)
	assert.IsType(t, model{}, m)
	cmd := m.Init()
	_, ok := cmd().(tea.BatchMsg)
	assert.True(t, ok)

	assert.Equal(t, m.state.scoreMode, 0)
	assert.Equal(t, m.help, help.New())
	assert.Equal(t, m.state.view, 0)

	for i, v := range []string{
		"Name",
		"Art/Animation",
		"Character/Cast",
		"Plot",
		"Bias",
		"Comments",
	} {
		assert.Equal(t, v, m.form.inputs[i].Placeholder)
	}
}

func TestDefaultButton(t *testing.T) {
	m := helperInitilizeModel(t)
	nm, cmd := m.buttonCommands()
	assert.Nil(t, cmd)

	switch nmt := nm.(type) {
	case model:
		assert.Equal(t, m.form.focusIndex+1, nmt.form.focusIndex)
	default:
		t.Errorf("incorrect model recieved. %T", nm)
	}
}

func TestSaveButton(t *testing.T) {
	m := helperInitilizeModel(t)
	mockValue(t, m)
	m.calculateScore()
	m.form.focusIndex = len(m.form.inputs)

	_, cmd := m.buttonCommands()
	assert.Nil(t, cmd)

	v, err := clipboard.ReadAll()
	assert.NoError(t, err)
	assert.Equal(t, "7.2", v)
}

func TestResetButton(t *testing.T) {
	m := helperInitilizeModel(t)
	mockValue(t, m)
	m.calculateScore()
	m.form.focusIndex = len(m.form.inputs) + 1
	nm, cmd := m.buttonCommands()

	assert.Nil(t, cmd)

	switch nm := nm.(type) {
	case model:
		assert.Equal(t, 0, nm.form.focusIndex)
		assert.Equal(t, float32(0.0), nm.state.score)
	default:
		t.Errorf("incorrect model recieved. %T", nm)
	}

	t.Run("test quit button", func(t *testing.T) {
		m := helperInitilizeModel(t)
		m.form.focusIndex = len(m.form.inputs) + 2
		_, cmd := m.buttonCommands()
		_, ok := cmd().(tea.QuitMsg)
		assert.True(t, ok)
	})
}
