package morse

import (
	"strings"
)

//ToText converts a morse string to his textual rapresentation
func ToText(morse string) string {
	c := DefaultConverter

	out := make([]rune, 0, int(float64(len(morse))/averageSize))

	words := strings.Split(morse, c.charSeparator+Space+c.charSeparator)
	for _, word := range words {
		chars := strings.Split(word, c.charSeparator)

		for _, ch := range chars {
			text, ok := c.morseToRune[ch]
			if !ok {
				out = append(out, []rune(c.Handling(ErrNoEncoding{string(text)}))...)
				continue
			}
			out = append(out, text)
		}
		out = append(out, ' ')
	}

	if len(words) > 0 {
		out = out[:len(out)-1]
	}

	return string(out)
}

//ToMorse converts a text to his morse rapresentation
func ToMorse(text string) string {
	c := DefaultConverter

	out := make([]rune, 0, int(float64(len(text))*averageSize))

	for _, ch := range text {
		out = append(out, []rune(c.runeToMorse[ch])...)
		out = append(out, []rune(c.charSeparator)...)
	}

	//Remove last charSeparator
	if len(text) > 0 {
		out = out[:len(out)-len(c.charSeparator)]
	}

	return string(out)
}
