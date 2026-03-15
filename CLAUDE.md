# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

kbd-cli is a terminal-based monkeytype-style typing speed test written in Go, featuring live keyboard visualization with key flash feedback.

## Build & Run

```bash
go build          # produces ./kbd-cli binary
./kbd-cli         # run the app
```

There are no tests, linter config, or Makefile currently.

## Architecture

Single `main` package using the **Bubble Tea v2** MVC framework (bubbletea + lipgloss).

**State machine with 3 phases:** `phaseReady` (duration selection) → `phaseTyping` (active test with countdown) → `phaseDone` (WPM/accuracy results).

**File responsibilities:**
- `main.go` — entry point, initializes Bubble Tea program
- `model.go` — core state (`model` struct), `Init`/`Update`/`View` methods, game logic, WPM/accuracy calculations
- `render.go` — view rendering for each phase (ready/typing/done screens)
- `keyboard.go` — QWERTY keyboard layout definitions (`KeyDef`) and key state tracking (`KeyState`)
- `styles.go` — lipgloss terminal styles
- `words.go` — 200-word list and random word generation

**Key types:** `model` (game state), `phase` (state enum), `KeyDef` (key layout), `KeyState` (untested/flashing).

**Timer system:** `tickMsg` for countdown, `flashTickMsg` for 150ms key flash animations.
