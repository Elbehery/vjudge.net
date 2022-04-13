package main

import "testing"

func TestGreeting(t *testing.T) {
	output := run()
	if output != str {
		t.Fatalf("incorrect greeting, expected %v, but got %v instead", str, output)
	}
}
