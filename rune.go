package morse

import (
	"unicode"
)

//RuneToMorse return the morse rapresentation of the rune
//If the rune is not a recognized morse character RuneToMorse will return an empty string
//Lowercase runes are converted to uppercase
//
//For Example: 'T' -> "-"
func RuneToMorse(ch rune) string {
	ch = unicode.ToUpper(ch)
	return DefaultMorse[ch]
}

//RuneToText return the character rapresented by the input string
//If the string is not recognizedc as a morse sequenze RuneToText will return a null rune
//
//For Example: "-" -> 'T'
func RuneToText(char string) rune {
	return reverseDefaultMorse[char]
}
