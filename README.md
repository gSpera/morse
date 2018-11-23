Morse
=====
[![Documentation](https://godoc.org/github.com/gSpera/morse?status.svg)](http://godoc.org/github.com/gSpera/morse)
[![Go Report Card](https://goreportcard.com/badge/github.com/gSpera/morse)](https://goreportcard.com/report/github.com/gSpera/morse)

Morse is a simple library for encoding and decoding between morse and text.

Support
=======
This library supports the default morse (as defined by **ITU-R M.1677-1**) code but custom ones can be used freely using a custom [EncodingMap](https://godoc.org/github.com/gSpera/morse#EncodingMap)

Tool
====
You can find a simple tool in the [cmd/morse](cmd/morse) directory
This tool can be used for converting to/from morse
```bash
$ morse > out.morse
test
this is morse.
^C
$ cat out.morse
- . ... - .-.-- .... .. ...   .. ...   -- --- .-. ... . .-.-.- .-.-
$ morse -D < out.morse
TEST
THIS IS MORSE.
```
For more uses look use `--help`

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
You can see more examples on the [godoc documentation](https://godoc.org/github.com/gSpera/morse)
