package morse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConverterOption(t *testing.T) {
	tm := []struct {
		name   string
		input  Converter
		fn     []ConverterOption
		output Converter
	}{
		{"No Functions", DefaultConverter, []ConverterOption{}, DefaultConverter},
		{"WithLowercaseHandling", Converter{}, []ConverterOption{WithLowercaseHandling(true)}, Converter{convertToUpper: true}},
		{"WithHandler", Converter{}, []ConverterOption{WithHandler(IgnoreHandler)}, Converter{Handling: IgnoreHandler}},
		{"WithCharSeparator", Converter{}, []ConverterOption{WithCharSeparator("separator")}, Converter{charSeparator: "separator"}},
		{"WithWordSeparator", Converter{}, []ConverterOption{WithWordSeparator("separator")}, Converter{wordSeparator: "separator"}},
	}

	for _, tt := range tm {
		t.Run(tt.name, func(t *testing.T) {
			in := tt.input
			for _, fn := range tt.fn {
				in = fn(in)
			}

			if err := compareConverter(tt.output, in); err != nil {
				t.Fatalf("Converter are not the same: %v", err)
			}
		})
	}
}

func compareConverter(a, b Converter) error {
	if a.charSeparator != b.charSeparator {
		return fmt.Errorf("Char Separetor are different: %q != %q", a.charSeparator, b.charSeparator)
	}
	if a.convertToUpper != b.convertToUpper {
		return fmt.Errorf("Convert to Upper are different: %t != %t", a.convertToUpper, b.convertToUpper)
	}
	if (a.Handling == nil && b.Handling != nil) ||
		(b.Handling == nil && a.Handling != nil) {
		return fmt.Errorf("Handlers are different")
	}
	if !reflect.DeepEqual(a.morseToRune, b.morseToRune) {
		return fmt.Errorf("MorseToRune are different")
	}
	if !reflect.DeepEqual(a.runeToMorse, b.runeToMorse) {
		return fmt.Errorf("RuneToMorse are different")
	}

	return nil
}
