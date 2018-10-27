package morse

import "fmt"

//ErrNoEncoding is the error used when there is no rapresentation
type ErrNoEncoding struct{ Text string }

func (e ErrNoEncoding) Error() string { return fmt.Sprintf("No encoding for: %q", e.Text) }

const averageSize = 4.53 //Magic

//Morse letters definitions
const (
	A = ".-"
	B = "-..."
	C = "-.-."
	D = "-.."
	E = "."
	F = "..-."
	G = "--."
	H = "...."
	I = ".."
	J = ".---"
	K = "-.-"
	L = ".-.."
	M = "--"
	N = "-."
	O = "---"
	P = ".--."
	Q = "--.-"
	R = ".-."
	S = "..."
	T = "-"
	U = "..-"
	V = "...-"
	W = ".--"
	X = "-..-"
	Y = "-.--"
	Z = "--.."

	Space = " "
)

var runeMap = map[rune]string{
	'A': A,
	'B': B,
	'C': C,
	'D': D,
	'E': E,
	'F': F,
	'G': G,
	'H': H,
	'I': I,
	'J': J,
	'K': K,
	'L': L,
	'M': M,
	'N': N,
	'O': O,
	'P': P,
	'Q': Q,
	'R': R,
	'S': S,
	'T': T,
	'U': U,
	'V': V,
	'W': W,
	'X': X,
	'Y': Y,
	'Z': Z,
	' ': Space,
}

var stringMap = map[string]rune{
	A:     'A',
	B:     'B',
	C:     'C',
	D:     'D',
	E:     'E',
	F:     'F',
	G:     'G',
	H:     'H',
	I:     'I',
	J:     'J',
	K:     'K',
	L:     'L',
	M:     'M',
	N:     'N',
	O:     'O',
	P:     'P',
	Q:     'Q',
	R:     'R',
	S:     'S',
	T:     'T',
	U:     'U',
	V:     'V',
	W:     'W',
	X:     'X',
	Y:     'Y',
	Z:     'Z',
	Space: ' ',
}

//IgnoreHandler ignores the error and returns nothing
func IgnoreHandler(error) string { return "" }

//PanicHandler is an handler that panics when an error occuours
func PanicHandler(err error) string { panic(err) }

//DefaultConverter is the default converter, it uses the exported morse set and has an IgnoreHandler, the separation character is a space
var DefaultConverter = Converter{
	charSeparator: " ",
	morseToRune:   stringMap,
	runeToMorse:   runeMap,

	Handling: IgnoreHandler,
}
