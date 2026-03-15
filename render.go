package main

import (
	"fmt"
	"math"
	"strings"

	"charm.land/lipgloss/v2"
)

const wordsPerLine = 12

func (m model) renderView() string {
	var content string
	switch m.phase {
	case phaseReady:
		content = m.renderReady()
	case phaseTyping:
		content = m.renderTyping()
	case phaseDone:
		content = m.renderDone()
	}

	if m.width > 0 && m.height > 0 {
		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)
	}
	return content
}

func (m model) renderReady() string {
	title := titleStyle.Render("⌨ Typing Test")

	var durs []string
	for i, d := range durations {
		label := fmt.Sprintf("%ds", d)
		if i == m.durationIdx {
			durs = append(durs, selectedDurationStyle.Render(label))
		} else {
			durs = append(durs, unselectedDurationStyle.Render(label))
		}
	}
	selector := strings.Join(durs, subtleStyle.Render("  /  "))

	var wlParts []string
	for i, wl := range wordLists {
		if i == m.wordListIdx {
			wlParts = append(wlParts, selectedDurationStyle.Render(wl.name))
		} else {
			wlParts = append(wlParts, unselectedDurationStyle.Render(wl.name))
		}
	}
	wordListSelector := strings.Join(wlParts, subtleStyle.Render("  /  "))

	hint := subtleStyle.Render("← → duration • ↑ ↓ word list • start typing to begin • esc to reset")

	preview := m.renderWordLines(3)
	keyboard := m.renderKeyboard()

	return lipgloss.JoinVertical(lipgloss.Center,
		"",
		title,
		"",
		selector,
		wordListSelector,
		"",
		preview,
		"",
		keyboard,
		"",
		hint,
		"",
	)
}

func (m model) renderTyping() string {
	timer := timerStyle.Render(fmt.Sprintf("%d", m.timeLeft))
	text := m.renderWordLines(3)
	keyboard := m.renderKeyboard()

	return lipgloss.JoinVertical(lipgloss.Center,
		"",
		timer,
		"",
		text,
		"",
		keyboard,
		"",
	)
}

func (m model) renderDone() string {
	title := titleStyle.Render("⌨ Results")

	wpm := resultValueStyle.Render(fmt.Sprintf("%.0f", m.wpm()))
	acc := resultValueStyle.Render(fmt.Sprintf("%.0f%%", m.accuracy()))

	wpmLine := resultLabelStyle.Render("wpm  ") + wpm
	accLine := resultLabelStyle.Render("acc  ") + acc

	hint := subtleStyle.Render("tab to restart • esc to menu • ctrl+c to exit")

	return lipgloss.JoinVertical(lipgloss.Center,
		"",
		title,
		"",
		wpmLine,
		accLine,
		"",
		hint,
		"",
	)
}

func (m model) renderWordLines(numLines int) string {
	currentLine := m.wordIdx / wordsPerLine

	firstLine := currentLine - numLines/2
	if firstLine < 0 {
		firstLine = 0
	}

	var lines []string
	totalWords := len(m.words)

	for line := firstLine; line < firstLine+numLines; line++ {
		startIdx := line * wordsPerLine
		if startIdx >= totalWords {
			break
		}
		endIdx := startIdx + wordsPerLine
		if endIdx > totalWords {
			endIdx = totalWords
		}

		var wordParts []string
		for i := startIdx; i < endIdx; i++ {
			wordParts = append(wordParts, m.renderWord(i))
		}
		lines = append(lines, strings.Join(wordParts, " "))
	}

	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}

func (m model) renderWord(idx int) string {
	target := []rune(m.words[idx])
	typed := m.input[idx]

	var result strings.Builder

	for i, ch := range target {
		if idx == m.wordIdx && i == len(typed) {
			result.WriteString(cursorStyle.Render(string(ch)))
		} else if i < len(typed) {
			if typed[i] == ch {
				result.WriteString(correctStyle.Render(string(ch)))
			} else {
				result.WriteString(incorrectStyle.Render(string(ch)))
			}
		} else {
			result.WriteString(untypedStyle.Render(string(ch)))
		}
	}

	if len(typed) > len(target) {
		for _, ch := range typed[len(target):] {
			result.WriteString(extraStyle.Render(string(ch)))
		}
	}

	return result.String()
}

// ── Keyboard rendering ──

func computeRowWidths(row []KeyDef) []int {
	widths := make([]int, len(row))
	cumUnits := 0.0
	cumChars := 0
	for i, k := range row {
		cumUnits += k.Width
		targetChars := int(math.Round(cumUnits * float64(keyUnit)))
		widths[i] = targetChars - cumChars
		cumChars = targetChars
	}
	return widths
}

func (m model) renderKbdKey(k KeyDef, totalWidth int) string {
	if k.Spacer {
		return strings.Repeat(" ", totalWidth)
	}

	style := kbdUntestedStyle
	if m.keyStates[k.ID] == KeyFlashing {
		style = kbdFlashingStyle
	}

	innerWidth := totalWidth - 2
	if innerWidth < 1 {
		innerWidth = 1
	}
	return style.Width(innerWidth).Render(k.Label)
}

func (m model) renderKbdRow(row []KeyDef) string {
	widths := computeRowWidths(row)
	keys := make([]string, len(row))
	for i, k := range row {
		keys[i] = m.renderKbdKey(k, widths[i])
	}
	return lipgloss.JoinHorizontal(lipgloss.Center, keys...)
}

func (m model) renderKeyboard() string {
	rows := make([]string, len(m.kbdLayout))
	for i, row := range m.kbdLayout {
		rows[i] = m.renderKbdRow(row)
	}
	return lipgloss.JoinVertical(lipgloss.Center, rows...)
}
