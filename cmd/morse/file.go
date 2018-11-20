package main

import (
	"fmt"
	"io"
	"os"
)

//PathFlag is a special flag that need that handles the checking of file existence and the stdin name "-"
type PathFlag string

//String implement flag.Value
func (p *PathFlag) String() string { return string(*p) }

//Set implement flag.Value, it returns an error if the file doesn't exist
func (p *PathFlag) Set(filepath string) error {
	if filepath == "-" { //Stdin/Stdout
		*p = PathFlag(filepath)
		return nil
	}

	_, err := os.Stat(filepath)
	switch {
	case os.IsNotExist(err):
		return fmt.Errorf("Cannot find file: %s", filepath)
	case err != nil:
		return err
	}

	*p = PathFlag(filepath)
	return nil
}

//Stream return a Read stream to the file.
func (p *PathFlag) Stream() (io.Reader, error) {
	if *p == "-" {
		return os.Stdin, nil
	}

	return os.Open(string(*p))
}
