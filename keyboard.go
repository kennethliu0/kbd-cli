package main

// KeyState represents the visual state of a key.
type KeyState int

const (
	KeyUntested KeyState = iota
	KeyFlashing
)

// KeyDef defines a single key on the keyboard layout.
type KeyDef struct {
	ID     string
	Label  string
	Width  float64
	Spacer bool
}

func spacer(units float64) KeyDef {
	return KeyDef{Spacer: true, Width: units}
}

func kbdKey(id, label string, width float64) KeyDef {
	return KeyDef{ID: id, Label: label, Width: width}
}

func buildKeyboardLayout() [][]KeyDef {
	return [][]KeyDef{
		{
			kbdKey("q", "Q", 1), kbdKey("w", "W", 1), kbdKey("e", "E", 1),
			kbdKey("r", "R", 1), kbdKey("t", "T", 1), kbdKey("y", "Y", 1),
			kbdKey("u", "U", 1), kbdKey("i", "I", 1), kbdKey("o", "O", 1),
			kbdKey("p", "P", 1),
		},
		{
			spacer(0.25),
			kbdKey("a", "A", 1), kbdKey("s", "S", 1), kbdKey("d", "D", 1),
			kbdKey("f", "F", 1), kbdKey("g", "G", 1), kbdKey("h", "H", 1),
			kbdKey("j", "J", 1), kbdKey("k", "K", 1), kbdKey("l", "L", 1),
		},
		{
			spacer(0.75),
			kbdKey("z", "Z", 1), kbdKey("x", "X", 1), kbdKey("c", "C", 1),
			kbdKey("v", "V", 1), kbdKey("b", "B", 1), kbdKey("n", "N", 1),
			kbdKey("m", "M", 1),
		},
		{
			spacer(2),
			kbdKey("space", "━━━", 4),
			spacer(2),
		},
	}
}
