package main

import "testing"

func TestGreeting(t *testing.T) {
	if Greeting() == "" {
		t.Fatal("greeting must not be empty")
	}
}
