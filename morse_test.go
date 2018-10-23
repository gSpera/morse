package morse_test

import (
	"testing"

	"github.com/gSpera/morse"
)

func TestRuneToMorse(t *testing.T) {
	tm := []struct {
		name   string
		input  rune
		output string
	}{
		{"Simple A", 'A', morse.A},
		{"Hardcoded A", 'A', ".-"},
		{"Non supported rune", '-', ""},
	}

	for _, tt := range tm {
		t.Run(tt.name, func(t *testing.T) {
			get := morse.ConvertRune(tt.input)
			if get != tt.output {
				t.Errorf("Expected [%s], got: [%s]", tt.output, get)
			}
		})
	}
}

func TestConverter_MorseToText(t *testing.T) {
	tm := []struct {
		name      string
		converter morse.Converter
		input     string
		output    string
	}{
		{
			"Simple Text",
			morse.DefaultConverter,
			".-.. --- .-. . --",
			"LOREM",
		},
		{
			"Empty String",
			morse.DefaultConverter,
			"",
			"",
		},
	}

	for _, tt := range tm {
		t.Run(tt.name, func(t *testing.T) {
			get := tt.converter.MorseToText(tt.input)
			if get != tt.output {
				t.Errorf("Expected [%s], got: [%s]", tt.output, get)
			}
		})
	}
}

func TestConverter_TextToMorse(t *testing.T) {
	tm := []struct {
		name      string
		converter morse.Converter
		input     string
		output    string
	}{
		{
			"Simple Text",
			morse.DefaultConverter,
			"LOREM",
			".-.. --- .-. . --",
		},
		{
			"Empty String",
			morse.DefaultConverter,
			"",
			"",
		},
	}

	for _, tt := range tm {
		t.Run(tt.name, func(t *testing.T) {
			get := tt.converter.TextToMorse(tt.input)
			if get != tt.output {
				t.Errorf("Expected [%s], got: [%s]", tt.output, get)
			}
		})
	}
}
