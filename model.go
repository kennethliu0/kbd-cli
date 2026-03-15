package main

import (
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
)

type phase int

const (
	phaseReady phase = iota
	phaseTyping
	phaseDone
)

var durations = []int{15, 30, 60}

type tickMsg time.Time
type flashTickMsg struct{ keyID string }

type model struct {
	phase        phase
	words        []string
	input        [][]rune
	wordIdx      int
	durationIdx  int
	wordListIdx  int
	timeLeft     int
	startTime    time.Time
	totalChars   int
	totalTyped   int
	correctChars int

	// Keyboard visualization
	kbdLayout [][]KeyDef
	keyStates map[string]KeyState

	// Terminal dimensions
	width, height int
}

func initialModel() model {
	return newTest(1, 0)
}

func newTest(durationIdx, wordListIdx int) model {
	words := generateWords(200, wordListIdx)
	input := make([][]rune, len(words))
	return model{
		phase:       phaseReady,
		words:       words,
		input:       input,
		durationIdx: durationIdx,
		wordListIdx: wordListIdx,
		timeLeft:    durations[durationIdx],
		kbdLayout:   buildKeyboardLayout(),
		keyStates:   make(map[string]KeyState),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func timerTick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func flashKey(keyID string) tea.Cmd {
	return tea.Tick(150*time.Millisecond, func(t time.Time) tea.Msg {
		return flashTickMsg{keyID: keyID}
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		if msg.Mod&tea.ModCtrl != 0 && msg.Code == 'c' {
			return m, tea.Quit
		}
		if msg.String() == "esc" {
			nm := newTest(m.durationIdx, m.wordListIdx)
			nm.width = m.width
			nm.height = m.height
			return nm, nil
		}
		return m.handleKey(msg)

	case tickMsg:
		if m.phase != phaseTyping {
			return m, nil
		}
		m.timeLeft--
		if m.timeLeft <= 0 {
			m.phase = phaseDone
			m.calcResults()
			return m, nil
		}
		return m, timerTick()

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case flashTickMsg:
		if m.keyStates[msg.keyID] == KeyFlashing {
			m.keyStates[msg.keyID] = KeyUntested
		}
		return m, nil
	}

	return m, nil
}

func (m *model) flashPressed(key string) tea.Cmd {
	// Map the typed character to a keyboard key ID
	id := strings.ToLower(key)
	if key == " " {
		id = "space"
	}
	if len(id) == 1 && id[0] >= 'a' && id[0] <= 'z' {
		m.keyStates[id] = KeyFlashing
		return flashKey(id)
	}
	if id == "space" {
		m.keyStates[id] = KeyFlashing
		return flashKey(id)
	}
	return nil
}

func (m model) handleKey(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch m.phase {
	case phaseReady:
		return m.handleReady(msg)
	case phaseTyping:
		return m.handleTyping(msg)
	case phaseDone:
		return m.handleDone(msg)
	}
	return m, nil
}

func (m model) handleReady(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "left":
		if m.durationIdx > 0 {
			m.durationIdx--
			m.timeLeft = durations[m.durationIdx]
		}
	case "right":
		if m.durationIdx < len(durations)-1 {
			m.durationIdx++
			m.timeLeft = durations[m.durationIdx]
		}
	case "up":
		if m.wordListIdx > 0 {
			nm := newTest(m.durationIdx, m.wordListIdx-1)
			nm.width = m.width
			nm.height = m.height
			return nm, nil
		}
	case "down":
		if m.wordListIdx < len(wordLists)-1 {
			nm := newTest(m.durationIdx, m.wordListIdx+1)
			nm.width = m.width
			nm.height = m.height
			return nm, nil
		}
	default:
		key := msg.String()
		if key == "space" || (len(key) == 1 && key[0] >= 32 && key[0] <= 126) {
			m.phase = phaseTyping
			m.startTime = time.Now()
			m.timeLeft = durations[m.durationIdx]
			m2, cmd := m.handleTyping(msg)
			return m2, tea.Batch(timerTick(), cmd)
		}
	}
	return m, nil
}

func (m model) handleTyping(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	var flashCmd tea.Cmd

	switch msg.String() {
	case "backspace":
		if len(m.input[m.wordIdx]) > 0 {
			m.input[m.wordIdx] = m.input[m.wordIdx][:len(m.input[m.wordIdx])-1]
		} else if m.wordIdx > 0 {
			m.wordIdx--
		}
	case "space", " ":
		if m.wordIdx < len(m.words)-1 {
			m.wordIdx++
		}
		flashCmd = m.flashPressed(" ")
	default:
		key := msg.String()
		if len(key) == 1 && key[0] >= 32 && key[0] <= 126 {
			m.input[m.wordIdx] = append(m.input[m.wordIdx], rune(key[0]))
			m.totalTyped++
			flashCmd = m.flashPressed(key)
		}
	}
	return m, flashCmd
}

func (m model) handleDone(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "tab":
		nm := newTest(m.durationIdx, m.wordListIdx)
		nm.width = m.width
		nm.height = m.height
		return nm, nil
	}
	return m, nil
}

func (m *model) calcResults() {
	m.correctChars = 0
	m.totalTyped = 0
	for i, word := range m.words {
		typed := m.input[i]
		target := []rune(word)
		m.totalTyped += len(typed)
		if i < m.wordIdx || (i == m.wordIdx && len(typed) > 0) {
			for j, ch := range typed {
				if j < len(target) && ch == target[j] {
					m.correctChars++
				}
			}
			if i < m.wordIdx {
				m.correctChars++
				m.totalTyped++
			}
		}
	}
	m.totalChars = m.correctChars
}

func (m model) wpm() float64 {
	elapsed := durations[m.durationIdx]
	minutes := float64(elapsed) / 60.0
	return float64(m.totalChars) / 5.0 / minutes
}

func (m model) accuracy() float64 {
	if m.totalTyped == 0 {
		return 100.0
	}
	return float64(m.correctChars) / float64(m.totalTyped) * 100.0
}

func (m model) View() tea.View {
	v := tea.NewView(m.renderView())
	v.AltScreen = true
	v.WindowTitle = "kbd-cli"
	return v
}
