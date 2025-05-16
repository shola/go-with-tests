package main

import "testing"

// The TDD process and why the steps are important:
// 1) Write a failing test and see it fail so we know we have written a relevant
// test for our requirements and seen that it produces an easy to understand
// description of the failure

// 2) Writing the smallest amount of code to make it pass so we know we have working
// software

// 3) Then refactor, backed with the safety of our tests to ensure we have
// well-crafted code that is easy to work with

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Jacque", "French")
		want := "Bonjour, Jacque"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
