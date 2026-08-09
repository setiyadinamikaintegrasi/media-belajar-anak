package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestGreeting(t *testing.T) {
	if got := Greeting(); got != "Media Belajar Anak API" {
		t.Fatalf("Greeting() = %q", got)
	}
}

func TestRunWritesGreeting(t *testing.T) {
	var output bytes.Buffer
	if err := run(&output); err != nil {
		t.Fatalf("run() error = %v", err)
	}
	if got := output.String(); got != "Media Belajar Anak API\n" {
		t.Fatalf("run() output = %q", got)
	}
}

func TestMainWritesGreeting(t *testing.T) {
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("os.Pipe() error = %v", err)
	}
	original := os.Stdout
	os.Stdout = writer
	t.Cleanup(func() {
		os.Stdout = original
		reader.Close()
		writer.Close()
	})

	main()
	os.Stdout = original
	if err := writer.Close(); err != nil {
		t.Fatalf("writer.Close() error = %v", err)
	}
	output, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("io.ReadAll() error = %v", err)
	}
	if got := string(output); got != "Media Belajar Anak API\n" {
		t.Fatalf("main() output = %q", got)
	}
}
