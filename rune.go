package morse

import (
	"unicode"
)

//RuneToMorse return the morse representation of the rune
//If the rune is not a recognized morse character RuneToMorse will return an empty string
//Lowercase runes are converted to uppercase
//
//For Example: 'T' -> "-"
func RuneToMorse(ch rune) string {
	ch = unicode.ToUpper(ch)
	return DefaultMorse[ch]
}

//RuneToText return the character represented by the input string
//If the string is not recognized as a morse sequence RuneToText will return a null rune
//
//For Example: "-" -> 'T'
func RuneToText(char string) rune {
	return reverseDefaultMorse[char]
}
