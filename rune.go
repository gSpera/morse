package morse

//RuneToMorse return the morse rapresentation of the rune
//If the rune is not a recognized morse character RuneToMorse will return an empty string
func RuneToMorse(ch rune) string {
	return DefaultMorse[ch]
}

//RuneToText return the character rapresented by the input string
//If the string is not recognizedc as a morse sequenze RuneToText will return a null rune
func RuneToText(char string) rune {
	return reverseDefaultMorse[char]
}
