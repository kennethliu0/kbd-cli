package main

// KeyState represents the visual state of a key.
type KeyState int

const (
	StateUntested KeyState = iota
	StateFlashing
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

// buildLayout returns the keyboard layout: alpha keys + spacebar with
// correct QWERTY stagger offsets (leading spacers only).
func buildLayout() [][]KeyDef {
	return [][]KeyDef{
		// QWERTY row
		{
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
		},
		// Home row
		{
			key("a", "A", 1, "a", "A"),
			key("s", "S", 1, "s", "S"),
			key("d", "D", 1, "d", "D"),
			key("f", "F", 1, "f", "F"),
			key("g", "G", 1, "g", "G"),
			key("h", "H", 1, "h", "H"),
			key("j", "J", 1, "j", "J"),
			key("k", "K", 1, "k", "K"),
			key("l", "L", 1, "l", "L"),
		},
		// Bottom row
		{
			key("z", "Z", 1, "z", "Z"),
			key("x", "X", 1, "x", "X"),
			key("c", "C", 1, "c", "C"),
			key("v", "V", 1, "v", "V"),
			key("b", "B", 1, "b", "B"),
			key("n", "N", 1, "n", "N"),
			key("m", "M", 1, "m", "M"),
		},
		// Space row
		{
			key("space", "Space", 4, " "),
		},
	}
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

