package main

import "charm.land/lipgloss/v2"

const keyUnit = 5 // one key unit in terminal characters

var (
	untestedStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#1a1a1a")).
			Foreground(lipgloss.Color("#555555")).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#333333")).
			Align(lipgloss.Center).
			AlignVertical(lipgloss.Center)

	flashingStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#FFD700")).
			Foreground(lipgloss.Color("#000000")).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FFFFFF")).
			Align(lipgloss.Center).
			AlignVertical(lipgloss.Center).
			Bold(true)

	testedStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#2E8B57")).
			Foreground(lipgloss.Color("#FFFFFF")).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#3CB371")).
			Align(lipgloss.Center).
			AlignVertical(lipgloss.Center)

	undetectableStyle = lipgloss.NewStyle().
				Background(lipgloss.Color("#0a0a0a")).
				Foreground(lipgloss.Color("#333333")).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#1a1a1a")).
				Align(lipgloss.Center).
				AlignVertical(lipgloss.Center)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFD700")).
			Bold(true).
			Align(lipgloss.Center)

	statusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888")).
			Align(lipgloss.Center)

	progressTestedStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#2E8B57"))

	progressUntestedStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#333333"))
)

func styleForState(state KeyState) lipgloss.Style {
	switch state {
	case StateFlashing:
		return flashingStyle
	case StateTested:
		return testedStyle
	case StateUndetectable:
		return undetectableStyle
	default:
		return untestedStyle
	}
}

func keyCharWidth(units float64) int {
	w := int(units*keyUnit + 0.5)
	if w < 3 {
		w = 3
	}
	return w
}
