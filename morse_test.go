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

func TestToText(t *testing.T) {
	tm := []struct {
		name   string
		input  string
		output string
	}{
		{
			"Simple",
			"--..",
			"Z",
		},
		{
			"Empty",
			"",
			"",
		},
	}

	for _, tt := range tm {
		t.Run(tt.name, func(t *testing.T) {
			out := morse.ToText(tt.input)
			if out != tt.output {
				t.Errorf("Expected: %q; got: %q", tt.output, out)
			}
		})
	}
}

func TestHandler(t *testing.T) {
	t.Run("IgnoreHandler", func(t *testing.T) {
		conv := morse.DefaultConverter
		conv.Handling = morse.IgnoreHandler
		out := conv.MorseToText("--------")
		if out != "" {
			t.Errorf("Expected \"\", got: %q", out)
		}
	})
	t.Run("PanicHandler", func(t *testing.T) {
		defer func() {
			if out := recover(); out == nil {
				t.Error("Expected Panic")
			}
		}()

		conv := morse.DefaultConverter
		conv.Handling = morse.PanicHandler
		conv.MorseToText("-------")
	})
}

func TestErrors(t *testing.T) {
	t.Run("ErrNoEncoding", func(t *testing.T) {
		err := morse.ErrNoEncoding{Text: "Text"}
		out := err.Error()
		expected := "No encoding for: \"Text\""
		if out != expected {
			t.Errorf("Expected: %q; got: %q", expected, out)
		}
	})
}
