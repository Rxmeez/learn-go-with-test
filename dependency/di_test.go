package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Rameez")

	got := buffer.String()
	want := "Hello, Rameez"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
