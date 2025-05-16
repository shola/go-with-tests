package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Olu")
	want := "Hello, Olu"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
