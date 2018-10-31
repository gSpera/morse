package morse

import (
	"io"
	"strings"
)

//Converter is a Morse from/to Text converter, it handles the conversion and error handling
type Converter struct {
	runeToMorse   map[rune]string
	morseToRune   map[string]rune
	charSeparator string

	Handling ErrorHandler
}

//ErrorHandler is a function used by Converter when it encounter an unknown character
type ErrorHandler func(error) string

//ToText converts a morse string to his textual rapresentation
func (c Converter) ToText(morse string) string {
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
func (c Converter) ToMorse(text string) string {
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

//ToMorseWriter translate all the text written to the returned io.Writer in morse code and writes it in the input io.Writer
func (c Converter) ToMorseWriter(output io.Writer) io.Writer {
	return translateToMorse{conv: c, buffer: make([]byte, 10), output: output}
}

//ToTextWriter translate all the text written to the returned io.Writer from morse code and writes it in the input io.Writer
func (c Converter) ToTextWriter(output io.Writer) io.Writer {
	return translateToText{conv: c, buffer: make([]byte, 10), output: output}
}
