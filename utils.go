package morse

import (
	"io"
)

//ToText converts a morse string to his textual rapresentation
func ToText(morse string) string { return DefaultConverter.MorseToText(morse) }

//ToMorse converts a text to his morse rapresentation
func ToMorse(text string) string { return DefaultConverter.TextToMorse(text) }

//ToMorseWriter translate all the text written to the returned io.Writer in morse code and writes it in the input io.Writer
func ToMorseWriter(output io.Writer) io.Writer { return DefaultConverter.ToMorseWriter(output) }

//ToTextWriter translate all the text written to the returned io.Writer from morse code and writes it in the input io.Writer
func ToTextWriter(output io.Writer) io.Writer { return DefaultConverter.ToTextWriter(output) }

type translateToMorse struct {
	conv   Converter
	buffer []byte

	input  io.Reader
	output io.Writer
}

//Text -> Morse
func (t translateToMorse) Write(data []byte) (int, error) {
	morse := t.conv.TextToMorse(string(data))
	return t.output.Write([]byte(morse))
}

type translateToText struct {
	conv   Converter
	buffer []byte

	input  io.Reader
	output io.Writer
}

//Morse -> Text
func (t translateToText) Write(data []byte) (int, error) {
	morse := t.conv.MorseToText(string(data))
	return t.output.Write([]byte(morse))
}
