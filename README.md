Morse
=====
[![Documentation](https://godoc.org/github.com/gSpera/Morse?status.svg)](http://godoc.org/github.com/gSpera/Morse)
[![Go Report Card](https://goreportcard.com/badge/github.com/gSpera/Morse)](https://goreportcard.com/report/github.com/gSpera/morse)

Morse is a simple library for encoding and decoding between morse and text.

Support
=======
This library support the default morse code but custom ones can be used freely using a custom EncodingMap

Examples
========
```go
text := "MORSE IS AWESOME"

//Convert to morse
textInMorse := morse.ToMorse(text)
fmt.Println(textInMorse) //-- --- .-. ... .   .. ...   .- .-- . ... --- -- .

//Back to text
backToText := morse.ToText(textInMorse)
fmt.Println(backToText) //MORSE IS AWESOME
```
You can see more examples on the [godoc documentation](https://godoc.org/github.com/gSpera/Morse)
