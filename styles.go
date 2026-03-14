package main

import "charm.land/lipgloss/v2"

const keyUnit = 7 // one key unit in terminal characters

var (
	untestedStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#1a1a1a")).
			Foreground(lipgloss.Color("#555555")).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#333333")).
			Align(lipgloss.Center)

	flashingStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#FFD700")).
			Foreground(lipgloss.Color("#000000")).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FFFFFF")).
			Align(lipgloss.Center).
			Bold(true)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFD700")).
			Bold(true).
			Align(lipgloss.Center)

	statusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888")).
			Align(lipgloss.Center)
)

func styleForState(state KeyState) lipgloss.Style {
	switch state {
	case StateFlashing:
		return flashingStyle
	default:
		return untestedStyle
	}
}

func keyCharWidth(units float64) int {
	w := int(units*keyUnit + 0.5)
	if w < 5 {
		w = 5
	}
	return w
}
