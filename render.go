package main

import (
	"math"
	"strings"

	"charm.land/lipgloss/v2"
)

// computeRowWidths converts fractional key-unit widths to integer char widths
// using cumulative rounding so every row with the same total units produces
// the same total char width.
func computeRowWidths(row []KeyDef) []int {
	widths := make([]int, len(row))
	cumUnits := 0.0
	cumChars := 0
	for i, k := range row {
		cumUnits += k.Width
		targetChars := int(math.Round(cumUnits * keyUnit))
		widths[i] = targetChars - cumChars
		cumChars = targetChars
	}
	return widths
}

func (m model) renderKeyWithWidth(k KeyDef, totalWidth int) string {
	if k.Spacer {
		return strings.Repeat(" ", totalWidth)
	}

	state, ok := m.keyStates[k.ID]
	if !ok {
		state = StateUntested
	}

	style := styleForState(state)
	// Subtract 2 for left+right border characters
	innerWidth := totalWidth - 2
	if innerWidth < 1 {
		innerWidth = 1
	}
	return style.Width(innerWidth).Render(k.Label)
}

func (m model) renderRow(row []KeyDef) string {
	widths := computeRowWidths(row)
	keys := make([]string, len(row))
	for i, k := range row {
		keys[i] = m.renderKeyWithWidth(k, widths[i])
	}
	return lipgloss.JoinHorizontal(lipgloss.Center, keys...)
}

func (m model) renderSection(section [][]KeyDef) string {
	rows := make([]string, len(section))
	for i, row := range section {
		rows[i] = m.renderRow(row)
	}
	return lipgloss.JoinVertical(lipgloss.Center, rows...)
}

func (m model) renderView() string {
	keyboard := m.renderSection(m.layout)

	title := titleStyle.Render("⌨ Keyboard Tester")
	footer := statusStyle.Render("Ctrl+C to exit")

	return lipgloss.JoinVertical(lipgloss.Center,
		"",
		title,
		"",
		keyboard,
		"",
		footer,
		"",
	)
}
