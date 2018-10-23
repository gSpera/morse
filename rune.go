package morse

//ConvertRune return the morse rapresentation of the rune
//If the rune is not a recognized morse character ConvertRune will return an empty string
func ConvertRune(ch rune) string {
	return runeMap[ch]
}

//ConvertMorseToRune return the character rapresented by the input string
//If the string is not recognizedc as a morse sequenze ConvertMorseToRune will return a null rune
func ConvertMorseToRune(char string) rune {
	return stringMap[char]
}
