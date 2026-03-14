package main

import "charm.land/lipgloss/v2"

var (
	// Text styles for the typing area
	untypedStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#555555"))
	correctStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))
	incorrectStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF4444"))
	cursorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700")).Underline(true)
	extraStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF4444")).Underline(true)

	// UI chrome
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFD700")).
			Bold(true)

	timerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFD700")).
			Bold(true)

	subtleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#555555"))

	resultLabelStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#888888"))

	resultValueStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFD700")).
				Bold(true)

	selectedDurationStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFD700")).
				Bold(true).
				Underline(true)

	unselectedDurationStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#555555"))

	// Keyboard visualization
	keyUnit = 7

	kbdUntestedStyle = lipgloss.NewStyle().
				Background(lipgloss.Color("#1a1a1a")).
				Foreground(lipgloss.Color("#555555")).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#333333")).
				Align(lipgloss.Center)

	kbdFlashingStyle = lipgloss.NewStyle().
				Background(lipgloss.Color("#FFD700")).
				Foreground(lipgloss.Color("#000000")).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#FFFFFF")).
				Align(lipgloss.Center).
				Bold(true)
)
