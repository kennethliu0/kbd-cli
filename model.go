package main

import (
	"time"

	tea "charm.land/bubbletea/v2"
)

type flashTickMsg struct{ keyID string }

type model struct {
	keyStates map[string]KeyState
	layout    [][]KeyDef
	keyLookup map[string]string
}

func initialModel() model {
	layout := buildLayout()
	lookup := buildKeyLookup(layout)

	return model{
		keyStates: make(map[string]KeyState),
		layout:    layout,
		keyLookup: lookup,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		// Ctrl+C is the only exit
		if msg.Mod&tea.ModCtrl != 0 && msg.Code == 'c' {
			return m, tea.Quit
		}

		var cmds []tea.Cmd

		// Match the key itself
		keyStr := msg.String()
		if keyID, ok := m.keyLookup[keyStr]; ok {
			cmds = append(cmds, m.markKey(keyID)...)
		}

		// Also try matching just the character for shifted combos
		if msg.Code > 0 {
			c := string(rune(msg.Code))
			if keyID, ok := m.keyLookup[c]; ok {
				cmds = append(cmds, m.markKey(keyID)...)
			}
		}

		return m, tea.Batch(cmds...)

	case flashTickMsg:
		if m.keyStates[msg.keyID] == StateFlashing {
			m.keyStates[msg.keyID] = StateUntested
		}
		return m, nil
	}

	return m, nil
}

func (m model) View() tea.View {
	v := tea.NewView(m.renderView())
	v.AltScreen = true
	return v
}

func (m *model) markKey(keyID string) []tea.Cmd {
	m.keyStates[keyID] = StateFlashing
	return []tea.Cmd{
		tea.Tick(150*time.Millisecond, func(t time.Time) tea.Msg {
			return flashTickMsg{keyID: keyID}
		}),
	}
}
