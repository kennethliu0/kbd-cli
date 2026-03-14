package main

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
)

func (m model) renderKey(k KeyDef) string {
	if k.Spacer {
		w := keyCharWidth(k.Width)
		return strings.Repeat(" ", w+2) // +2 for border space
	}

	state, ok := m.keyStates[k.ID]
	if !ok {
		if len(k.MatchIDs) == 0 {
			state = StateUndetectable
		} else {
			state = StateUntested
		}
	}

	style := styleForState(state)
	w := keyCharWidth(k.Width)
	// Subtract 2 for left+right border characters
	innerWidth := w - 2
	if innerWidth < 1 {
		innerWidth = 1
	}
	return style.Width(innerWidth).Render(k.Label)
}

func (m model) renderRow(row []KeyDef) string {
	keys := make([]string, len(row))
	for i, k := range row {
		keys[i] = m.renderKey(k)
	}
	return lipgloss.JoinHorizontal(lipgloss.Center, keys...)
}

func (m model) renderSection(section [][]KeyDef) string {
	rows := make([]string, len(section))
	for i, row := range section {
		rows[i] = m.renderRow(row)
	}
	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}

func (m model) renderProgressBar() string {
	total := m.totalKeys
	if total == 0 {
		total = 1
	}
	pct := m.testedKeys * 100 / total
	barWidth := 20
	filled := m.testedKeys * barWidth / total

	bar := progressTestedStyle.Render(strings.Repeat("█", filled)) +
		progressUntestedStyle.Render(strings.Repeat("░", barWidth-filled))

	return fmt.Sprintf("Tested: %d/%d keys (%d%%) %s", m.testedKeys, m.totalKeys, pct, bar)
}

func (m model) renderView() string {
	mainSection := m.renderSection(m.main)
	navSection := m.renderSection(m.nav)
	numpadSection := m.renderSection(m.numpad)

	gap := "  "
	keyboard := lipgloss.JoinHorizontal(lipgloss.Top,
		mainSection, gap, navSection, gap, numpadSection,
	)

	title := titleStyle.Render("⌨ Keyboard Tester")
	progress := statusStyle.Render(m.renderProgressBar())
	footer := statusStyle.Render("Ctrl+C to exit • Dimmed keys cannot be detected in terminal")

	return lipgloss.JoinVertical(lipgloss.Center,
		"",
		title,
		"",
		keyboard,
		"",
		progress,
		footer,
		"",
	)
}
