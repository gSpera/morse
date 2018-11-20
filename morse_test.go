package morse_test

import (
	"bytes"
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
		{"Lowercase", 'a', ".-"},
	}

	for _, tt := range tm {
		t.Run(tt.name, func(t *testing.T) {
			get := morse.RuneToMorse(tt.input)
			if get != tt.output {
				t.Errorf("Expected [%s], got: [%s]", tt.output, get)
			}
		})
	}
}

func TestConverter_ToText(t *testing.T) {
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
		{
			"Trailing Separator with Handler",
			morse.NewConverter(morse.DefaultMorse,
				morse.WithHandler(func(error) string { return "A" }),
				morse.WithTrailingSeparator(true),
			),
			"",
			"A  ",
		},
	}

	for _, tt := range tm {
		t.Run(tt.name, func(t *testing.T) {
			get := tt.converter.ToText(tt.input)
			if get != tt.output {
				t.Errorf("Expected [%s], got: [%s]", tt.output, get)
			}
		})
	}
}

func TestConverter_ToMorse(t *testing.T) {
	tm := []struct {
		name      string
		converter morse.Converter
		input     string
		output    string
		panics    interface{}
	}{
		{
			"Simple Text",
			morse.DefaultConverter,
			"LOREM",
			".-.. --- .-. . --",
			nil,
		},
		{
			"Empty String",
			morse.DefaultConverter,
			"",
			"",
			nil,
		},
		{
			"Character not supported",
			morse.NewConverter(morse.EncodingMap{}, morse.WithHandler(morse.PanicHandler)),
			"A",
			"",
			morse.ErrNoEncoding{Text: "A"},
		},
		{
			"Character not supported - Ignore",
			morse.NewConverter(morse.EncodingMap{}, morse.WithHandler(morse.IgnoreHandler)),
			"A",
			"",
			nil,
		},
		{
			"Trailing Separator with Handler",
			morse.NewConverter(morse.EncodingMap{},
				morse.WithHandler(func(error) string { return "A" }),
				morse.WithTrailingSeparator(true),
			),

			" ",
			"A ",
			nil,
		},
	}

	for _, tt := range tm {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				err := recover()
				if tt.panics != err {
					t.Errorf("Expected panic: %v; got: %v", tt.panics, err)
					if err != nil {
						panic(err)
					}
				}
			}()

			get := tt.converter.ToMorse(tt.input)
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
		out := conv.ToText("--------")
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
		conv.ToText("-------")
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

func TestConverter_ToMorseWriter(t *testing.T) {
	tm := []struct {
		name   string
		input  string
		output string
	}{
		{
			"Letter",
			"G",
			"--.",
		},
		{
			"Text",
			"TEXT",
			"- . -..- -",
		},
	}

	buffer := bytes.NewBufferString("")
	writer := morse.ToMorseWriter(buffer)
	for _, tt := range tm {
		t.Run(tt.name, func(t *testing.T) {
			buffer.Reset()
			writer.Write([]byte(tt.input))
			output := buffer.String()
			if output != tt.output {
				t.Errorf("Expected: %q; got: %q", tt.output, output)
			}
		})
	}
}

func TestConverter_ToTextWriter(t *testing.T) {
	tm := []struct {
		name   string
		input  string
		output string
		err    error
		n      int
	}{
		{
			"Letter",
			"--.",
			"G",
			nil,
			3,
		},
		{
			"Text",
			"- . -..- -",
			"TEXT",
			nil,
			10,
		},
	}

	buffer := bytes.NewBufferString("")
	writer := morse.ToTextWriter(buffer)
	for _, tt := range tm {
		t.Run(tt.name, func(t *testing.T) {
			buffer.Reset()
			n, err := writer.Write([]byte(tt.input))
			if tt.err != err {
				t.Errorf("Expected error: %v; got: %v", tt.err, err)
			}
			if tt.n != n {
				t.Errorf("Expected n: %v; got: %v", tt.n, n)
			}

			output := buffer.String()
			if output != tt.output {
				t.Errorf("Expected: %q; got: %q", tt.output, output)
			}
		})
	}
}

func TestConverter_CharSeparator(t *testing.T) {
	separator := "separator"
	c := morse.NewConverter(morse.DefaultMorse,
		morse.WithCharSeparator(separator),
		morse.WithHandler(morse.PanicHandler),
	)
	out := c.CharSeparator()

	if out != separator {
		t.Errorf("Expected: %q; got: %q", separator, out)
	}
}

func TestConverter_EncodingMap(t *testing.T) {
	expectedMap := morse.DefaultMorse

	c := morse.NewConverter(expectedMap,
		morse.WithHandler(morse.PanicHandler),
	)
	out := c.EncodingMap()

	for k := range expectedMap {
		if expectedMap[k] != out[k] {
			t.Errorf("Checking: %q: Expected: %q; got: %q", k, expectedMap[k], out[k])
		}
	}
}

func Test_NewConverter(t *testing.T) {
	t.Run("nil map", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Error("Expected")
			}
		}()
		morse.NewConverter(nil)
	})
}
