package main

import (
	"golang.org/x/example/hello/reverse"
	"testing"
)

func TestReverse(t *testing.T) {
	got := reverse.String("Hello, OTUS!")
	expected := "!SUTO ,olleH"

	if got != expected {
		t.Errorf("Reverse() = %q; want %q", got, expected)
	}
}
