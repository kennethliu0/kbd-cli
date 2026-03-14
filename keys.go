package main

// KeyState represents the visual state of a key.
type KeyState int

const (
	StateUntested KeyState = iota
	StateFlashing
	StateTested
	StateUndetectable
)

// KeyDef defines a single key on the keyboard layout.
type KeyDef struct {
	ID       string   // unique key id ("q", "f1", "lshift")
	Label    string   // display text ("Q", "F1", "LShift")
	Width    float64  // width in key units (1.0 = standard key)
	MatchIDs []string // all msg.String() values that trigger this key
	Spacer   bool     // true if this is a gap, not a real key
}

func spacer(units float64) KeyDef {
	return KeyDef{Spacer: true, Width: units}
}

func key(id, label string, width float64, matchIDs ...string) KeyDef {
	return KeyDef{ID: id, Label: label, Width: width, MatchIDs: matchIDs}
}

func undetectable(id, label string, width float64) KeyDef {
	return KeyDef{ID: id, Label: label, Width: width}
}

// buildLayout returns the full 108-key layout as rows within three sections.
// Each section is a slice of rows ([][]KeyDef).
// Sections: main, nav cluster, numpad.
func buildLayout() (main, nav, numpad [][]KeyDef) {
	// ── Main section ──
	main = [][]KeyDef{
		// Row 0: Esc, gap, F1-F4, gap, F5-F8, gap, F9-F12
		{
			key("esc", "Esc", 1, "escape"),
			spacer(1),
			key("f1", "F1", 1, "f1"),
			key("f2", "F2", 1, "f2"),
			key("f3", "F3", 1, "f3"),
			key("f4", "F4", 1, "f4"),
			spacer(0.5),
			key("f5", "F5", 1, "f5"),
			key("f6", "F6", 1, "f6"),
			key("f7", "F7", 1, "f7"),
			key("f8", "F8", 1, "f8"),
			spacer(0.5),
			key("f9", "F9", 1, "f9"),
			key("f10", "F10", 1, "f10"),
			key("f11", "F11", 1, "f11"),
			key("f12", "F12", 1, "f12"),
		},
		// Row 1: Number row
		{
			key("backtick", "`", 1, "`", "~"),
			key("1", "1", 1, "1", "!"),
			key("2", "2", 1, "2", "@"),
			key("3", "3", 1, "3", "#"),
			key("4", "4", 1, "4", "$"),
			key("5", "5", 1, "5", "%"),
			key("6", "6", 1, "6", "^"),
			key("7", "7", 1, "7", "&"),
			key("8", "8", 1, "8", "*"),
			key("9", "9", 1, "9", "("),
			key("0", "0", 1, "0", ")"),
			key("minus", "-", 1, "-", "_"),
			key("equal", "=", 1, "=", "+"),
			key("backspace", "Bksp", 2, "backspace"),
		},
		// Row 2: QWERTY
		{
			key("tab", "Tab", 1.5, "tab"),
			key("q", "Q", 1, "q", "Q"),
			key("w", "W", 1, "w", "W"),
			key("e", "E", 1, "e", "E"),
			key("r", "R", 1, "r", "R"),
			key("t", "T", 1, "t", "T"),
			key("y", "Y", 1, "y", "Y"),
			key("u", "U", 1, "u", "U"),
			key("i", "I", 1, "i", "I"),
			key("o", "O", 1, "o", "O"),
			key("p", "P", 1, "p", "P"),
			key("lbracket", "[", 1, "[", "{"),
			key("rbracket", "]", 1, "]", "}"),
			key("backslash", "\\", 1.5, "\\", "|"),
		},
		// Row 3: Home row
		{
			key("capslock", "Caps", 1.75, "caps_lock"),
			key("a", "A", 1, "a", "A"),
			key("s", "S", 1, "s", "S"),
			key("d", "D", 1, "d", "D"),
			key("f", "F", 1, "f", "F"),
			key("g", "G", 1, "g", "G"),
			key("h", "H", 1, "h", "H"),
			key("j", "J", 1, "j", "J"),
			key("k", "K", 1, "k", "K"),
			key("l", "L", 1, "l", "L"),
			key("semicolon", ";", 1, ";", ":"),
			key("quote", "'", 1, "'", "\""),
			key("enter", "Enter", 2.25, "enter"),
		},
		// Row 4: Bottom row
		{
			key("lshift", "LShift", 2.25, "shift"),
			key("z", "Z", 1, "z", "Z"),
			key("x", "X", 1, "x", "X"),
			key("c", "C", 1, "c", "C"),
			key("v", "V", 1, "v", "V"),
			key("b", "B", 1, "b", "B"),
			key("n", "N", 1, "n", "N"),
			key("m", "M", 1, "m", "M"),
			key("comma", ",", 1, ",", "<"),
			key("period", ".", 1, ".", ">"),
			key("slash", "/", 1, "/", "?"),
			key("rshift", "RShift", 2.75, "shift"),
		},
		// Row 5: Modifier row
		{
			key("lctrl", "Ctrl", 1.25, "ctrl"),
			undetectable("lsuper", "Super", 1.25),
			key("lalt", "Alt", 1.25, "alt"),
			key("space", "Space", 6.25, " "),
			key("ralt", "Alt", 1.25, "alt"),
			undetectable("rsuper", "Super", 1.25),
			undetectable("menu", "Menu", 1.25),
			key("rctrl", "Ctrl", 1.25, "ctrl"),
		},
	}

	// ── Nav cluster ──
	nav = [][]KeyDef{
		// Row 0: PrtSc, ScrLk, Pause
		{
			undetectable("prtsc", "Prt", 1),
			key("scrolllock", "Scr", 1, "scroll_lock"),
			key("pause", "Pau", 1, "pause"),
		},
		// Row 1: Ins, Home, PgUp
		{
			key("insert", "Ins", 1, "insert"),
			key("home", "Home", 1, "home"),
			key("pgup", "PgUp", 1, "pgup"),
		},
		// Row 2: Del, End, PgDn
		{
			key("delete", "Del", 1, "delete"),
			key("end", "End", 1, "end"),
			key("pgdown", "PgDn", 1, "pgdown"),
		},
		// Row 3: empty (spacer row for alignment)
		{
			spacer(3),
		},
		// Row 4: arrows top (up)
		{
			spacer(1),
			key("up", "↑", 1, "up"),
			spacer(1),
		},
		// Row 5: arrows bottom
		{
			key("left", "←", 1, "left"),
			key("down", "↓", 1, "down"),
			key("right", "→", 1, "right"),
		},
	}

	// ── Numpad ──
	numpad = [][]KeyDef{
		// Row 0: empty (aligns with F-key row)
		{
			spacer(4),
		},
		// Row 1: NumLk, /, *, -
		{
			key("numlock", "Num", 1, "num_lock"),
			key("numdiv", "/", 1, "kp_divide"),
			key("nummul", "*", 1, "kp_multiply"),
			key("numsub", "-", 1, "kp_subtract"),
		},
		// Row 2: 7, 8, 9, +
		{
			key("num7", "7", 1, "kp_7"),
			key("num8", "8", 1, "kp_8"),
			key("num9", "9", 1, "kp_9"),
			key("numadd", "+", 1, "kp_add"),
		},
		// Row 3: 4, 5, 6
		{
			key("num4", "4", 1, "kp_4"),
			key("num5", "5", 1, "kp_5"),
			key("num6", "6", 1, "kp_6"),
			spacer(1),
		},
		// Row 4: 1, 2, 3, Enter
		{
			key("num1", "1", 1, "kp_1"),
			key("num2", "2", 1, "kp_2"),
			key("num3", "3", 1, "kp_3"),
			key("numenter", "Ent", 1, "kp_enter"),
		},
		// Row 5: 0 (2u), ., (enter continues)
		{
			key("num0", "0", 2, "kp_0"),
			key("numdot", ".", 1, "kp_decimal"),
			spacer(1),
		},
	}

	return main, nav, numpad
}

// buildKeyLookup creates a map from msg.String() value to keyID.
func buildKeyLookup(sections ...[][]KeyDef) map[string]string {
	lookup := make(map[string]string)
	for _, section := range sections {
		for _, row := range section {
			for _, k := range row {
				if k.Spacer {
					continue
				}
				for _, mid := range k.MatchIDs {
					lookup[mid] = k.ID
				}
			}
		}
	}
	return lookup
}

// countTestableKeys returns the number of keys that have at least one MatchID.
func countTestableKeys(sections ...[][]KeyDef) int {
	seen := make(map[string]bool)
	count := 0
	for _, section := range sections {
		for _, row := range section {
			for _, k := range row {
				if k.Spacer || len(k.MatchIDs) == 0 {
					continue
				}
				if !seen[k.ID] {
					seen[k.ID] = true
					count++
				}
			}
		}
	}
	return count
}
