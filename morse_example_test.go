package morse_test

import (
	"fmt"

	"github.com/gSpera/morse"
)

func Example() {
	text := "MORSE IS AWESOME"
	textInMorse := morse.ToMorse(text)
	fmt.Println(textInMorse)
	backToText := morse.ToText(textInMorse)
	fmt.Println(backToText)
	//Output: -- --- .-. ... .   .. ...   .- .-- . ... --- -- .
	//MORSE IS AWESOME
}
func ExampleConvertRune() {
	ch := 'G'
	str := morse.ConvertRune(ch)

	fmt.Printf("The letter %c converts to: %s", ch, str)
	//Output: The letter G converts to: --.
}
func ExampleConvertMorseToRune() {
	str := "--."
	ch := morse.ConvertMorseToRune(str)

	fmt.Printf("The morse code %s converts to: %c", str, ch)
	//Output: The morse code --. converts to: G
}
