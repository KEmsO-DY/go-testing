package hello

import "testing"

func TestHello(t *testing.T) {
	if got, want := Hello("Hexlet"), "Hello, Hexlet!"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
