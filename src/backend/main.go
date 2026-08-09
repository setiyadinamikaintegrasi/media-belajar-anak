package main

import (
	"fmt"
	"io"
	"os"
)

func Greeting() string {
	return "Media Belajar Anak API"
}

func run(output io.Writer) error {
	_, err := fmt.Fprintln(output, Greeting())
	return err
}

func main() {
	if err := run(os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
