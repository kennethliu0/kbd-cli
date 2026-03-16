# kbd-cli

A monkeytype-style typing speed test for the terminal, built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) and [Lip Gloss](https://github.com/charmbracelet/lipgloss).

![kbd-cli screenshot](screenshot.png)

## Features

- Timed typing tests (15s / 30s / 60s)
- Live keyboard visualization with key flash on press
- WPM and accuracy results
- Multiple word lists: English 200, English 1k, English 5k

## Install

```
go install github.com/kennethliu0/kbd-cli@latest
```

Or build from source:

```
go build
./kbd-cli
```

## Controls

- **← →** change test duration
- **↑ ↓** change word list
- **Start typing** to begin the test
- **Tab** to restart after results
- **Ctrl+C** to exit

## Word Lists

| List | Words | Description |
|------|-------|-------------|
| English 200 | 200 | Most common English words (default) |
| English 1k | 1,000 | Common English words |
| English 5k | 5,000 | Expanded English vocabulary |
