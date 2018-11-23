package morse

import "fmt"

//ErrNoEncoding is the error used when there is no rapresentation
//It's primary use is inside Handlers
type ErrNoEncoding struct{ Text string }

//Error implements the error interface
func (e ErrNoEncoding) Error() string { return fmt.Sprintf("No encoding for: %q", e.Text) }

//EncodingMap contains the definitions for converting between two encoding
//It converts from a text rune (for example 'A') to it's morse rapresentation (for example ".-")
type EncodingMap map[rune]string

//averageSize is the average size of a morse char
const averageSize = 4.53 //Magic

//Morse letters and figures definitions
const (
	A         = ".-"
	B         = "-..."
	C         = "-.-."
	D         = "-.."
	E         = "."
	AccentedE = "..-.."
	F         = "..-."
	G         = "--."
	H         = "...."
	I         = ".."
	J         = ".---"
	K         = "-.-"
	L         = ".-.."
	M         = "--"
	N         = "-."
	O         = "---"
	P         = ".--."
	Q         = "--.-"
	R         = ".-."
	S         = "..."
	T         = "-"
	U         = "..-"
	V         = "...-"
	W         = ".--"
	X         = "-..-"
	Y         = "-.--"
	Z         = "--.."

	One   = ".----"
	Two   = "..---"
	Three = "...--"
	Four  = "....-"
	Five  = "....."
	Six   = "-...."
	Seven = "--..."
	Eight = "---.."
	Nine  = "----."
	Zero  = "-----"

	Period       = ".-.-.-" //.
	Comma        = "--..--" //,
	Colon        = "---..." //:
	QuestionMark = "..--.." //?
	Apostrophe   = ".----." //'
	Hyphen       = "-....-" //-
	Division     = "-..-."  ///
	LeftBracket  = "-.--."  //(
	RightBracket = "-.--.-" //)
	IvertedComma = ".-..-." //“ ”
	DoubleHyphen = "-...-"  //=
	Cross        = ".-.-."  //+
	CommercialAt = ".--.-." //@

	Understood           = "...-."
	Error                = "........"
	InvitationToTransmit = "-.-"
	Wait                 = ".-..."
	EndOfWork            = "...-.-"
	StartingSignal       = "-.-.-"

	Space = " "
)

//DefaultMorse is the default map used to convert between morse and text
//The map contians all the standard codes defined as costants but doesn't include commands like Understood and Error
//This map may remain constant.
var DefaultMorse = EncodingMap{
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

	'1': ".----",
	'2': "..---",
	'3': "...--",
	'4': "....-",
	'5': ".....",
	'6': "-....",
	'7': "--...",
	'8': "---..",
	'9': "----.",
	'0': "-----",

	'.':  ".-.-.-",
	',':  "--..--",
	':':  "---...",
	'?':  "..--..",
	'\'': ".----.",
	'-':  "-....-",
	'/':  "-..-.",
	'(':  "-.--.",
	')':  "-.--.-",
	'“':  ".-..-.",
	'=':  "-...-",
	'+':  ".-.-.",
	'@':  ".--.-.",
	' ':  Space,
}

var reverseDefaultMorse = reverseEncodingMap(DefaultMorse)

//IgnoreHandler ignores the error and returns nothing
func IgnoreHandler(error) string { return "" }

//PanicHandler is an handler that panics when an error occuours
func PanicHandler(err error) string { panic(err) }

//DefaultConverter is the default converter, it uses the exported morse set and has an IgnoreHandler, the separation character is a space
//Lowercase letter are encoded as upper ones. DefaultConverter uses explicitly IgnoreHandler and adds the trailoing separator
var DefaultConverter = NewConverter(
	DefaultMorse,

	WithCharSeparator(" "),
	WithWordSeparator("   "),
	WithLowercaseHandling(true),
	WithHandler(IgnoreHandler),
	WithTrailingSeparator(false),
)
