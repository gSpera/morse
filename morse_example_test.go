package morse_test

import (
	"fmt"

	"github.com/gSpera/morse"
)

func Example() {
	text := "MORSE IS AWESOME"

	//Convert to morse
	textInMorse := morse.ToMorse(text)
	fmt.Println(textInMorse)

	//Back to text
	backToText := morse.ToText(textInMorse)
	fmt.Println(backToText)
	//Output: -- --- .-. ... .   .. ...   .- .-- . ... --- -- .
	//MORSE IS AWESOME
}
func ExampleRuneToMorse() {
	ch := 'G'
	str := morse.RuneToMorse(ch)

	fmt.Printf("The letter %c converts to: %s", ch, str)
	//Output: The letter G converts to: --.
}
func ExampleRuneToText() {
	str := "--."
	ch := morse.RuneToText(str)

	fmt.Printf("The morse code %s converts to: %c", str, ch)
	//Output: The morse code --. converts to: G
}
