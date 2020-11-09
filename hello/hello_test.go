package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Rameez")
	want := "Hello, Rameez"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
