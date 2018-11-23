package morse

import (
	"io"
	"strings"
	"unicode"
)

//ErrorHandler is a function used by Converter when it encounter an unknown character
//Returns the text to inserted at the place of the unknown character
//This may not(but can if necessary) corrupt the output inserting invalid morse character
type ErrorHandler func(error) string

//Converter is a Morse from/to Text converter, it handles the conversion and error handling
type Converter struct {
	runeToMorse       map[rune]string
	morseToRune       map[string]rune
	charSeparator     string
	wordSeparator     string
	convertToUpper    bool
	trailingSeparator bool

	Handling ErrorHandler
}

//NewConverter creates a new converter with the specified configuration
//convertingMap is an EncodingMap, it contains how the characters will be translated, usually this is set to DefaultMorse
//but a custom one can be used. A nil convertingMap will panic.
func NewConverter(convertingMap EncodingMap, options ...ConverterOption) Converter {
	if convertingMap == nil {
		panic("Using a nil EncodingMap")
	}

	morseToRune := reverseEncodingMap(convertingMap)

	c := Converter{
		runeToMorse:       convertingMap,
		morseToRune:       morseToRune,
		charSeparator:     " ",
		wordSeparator:     "",
		convertToUpper:    false,
		trailingSeparator: false,

		Handling: IgnoreHandler,
	}

	for _, opt := range options {
		c = opt(c)
	}

	//Set wordSeparator as default
	if c.wordSeparator == "" {

		//Use custom space if avaible
		sp, ok := c.runeToMorse[' ']
		if !ok {
			//Fallback to the default Space
			sp = Space
		}
		c.wordSeparator = c.charSeparator + sp + c.charSeparator
	}

	return c
}

//ToText converts a morse string to his textual rapresentation
//
//For Example: "- . ... -" -> "TEST"
func (c Converter) ToText(morse string) string {
	out := make([]rune, 0, int(float64(len(morse))/averageSize))

	words := strings.Split(morse, c.charSeparator+Space+c.charSeparator)
	for _, word := range words {
		chars := strings.Split(word, c.charSeparator)

		for _, ch := range chars {
			text, ok := c.morseToRune[ch]
			if !ok {
				hand := []rune(c.Handling(ErrNoEncoding{string(ch)}))
				out = append(out, hand...)

				//Add a charSeparator is the len of the result is not zero
				if len(hand) != 0 {
					out = append(out, []rune(c.charSeparator)...)
				}
				continue
			}
			out = append(out, text)
		}

		out = append(out, ' ')
	}

	//Remove last charSeparator
	if !c.trailingSeparator && len(out) >= len(c.charSeparator) {
		out = out[:len(out)-len(c.charSeparator)]
	}

	return string(out)
}

//ToMorse converts a text to his morse rapresentation
//Lowercase characters are automatically converted to Uppercase
//
//For Example: "Test" -> "- . ... -"
func (c Converter) ToMorse(text string) string {
	out := make([]rune, 0, int(float64(len(text))*averageSize))

	for _, ch := range text {
		if c.convertToUpper {
			ch = unicode.ToUpper(ch)
		}

		if _, ok := c.runeToMorse[ch]; !ok {
			hand := []rune(c.Handling(ErrNoEncoding{string(ch)}))
			out = append(out, hand...)

			//Add a charSeparator is the len of the result is not zero
			if len(hand) != 0 {
				out = append(out, []rune(c.charSeparator)...)
			}
			continue
		}

		out = append(out, []rune(c.runeToMorse[ch])...)
		out = append(out, []rune(c.charSeparator)...)
	}

	//Remove last charSeparator
	if !c.trailingSeparator && len(out) >= len(c.charSeparator) {
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

//CharSeparator returns the charSeparator of the converter
func (c Converter) CharSeparator() string { return c.charSeparator }

//EncodingMap returns a copy of the EncodingMap inside the Converter,
//modifing the returned map will not change the internal one
func (c Converter) EncodingMap() EncodingMap {
	ret := make(EncodingMap, len(c.runeToMorse))

	for k, v := range c.runeToMorse {
		ret[k] = v
	}

	return ret
}
