package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/gSpera/morse"
)

func main() {
	decode := flag.Bool("D", false, "Decodes input (Morse -> Text)")
	in := PathFlag("-")
	var out string
	flag.Var(&in, "in", "The input file; default to stdin")
	flag.StringVar(&out, "out", "-", "The output file; default to stdout")
	flag.Parse()

	var writer io.Writer
	if out == "-" {
		writer = os.Stdout
	} else {
		var err error
		writer, err = os.OpenFile(out, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Cannot create output stream:", err)
			os.Exit(1)
		}
	}

	reader, err := in.Stream()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot create reader stream:", err)
		os.Exit(1)
	}

	encodingMap := morse.DefaultConverter.EncodingMap()
	encodingMap['\n'] = ".-.-"

	converter := morse.NewConverter(encodingMap,
		morse.WithLowercaseHandling(true),
		morse.WithHandler(morse.IgnoreHandler),
	)

	var stream io.Writer
	if *decode {
		stream = converter.ToTextWriter(writer)
	} else {
		stream = converter.ToMorseWriter(writer)
	}

	if _, err := io.Copy(stream, reader); err != nil {
		panic(err)
	}
}
