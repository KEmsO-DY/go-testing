package even

import "testing"

func TestIsEven(t *testing.T) {

	if got, want := IsEven(10), true; got != want {
		t.Errorf("Error on value: %v, expected %v", got, want)
	}
	if got, want := IsEven(0), true; got != want {
		t.Errorf("Error on value: %v, expected %v", got, want)
	}
	if got, want := IsEven(-4), true; got != want {
		t.Errorf("Error on value: %v, expected %v", got, want)
	}
	if got, want := IsEven(-7), false; got != want {
		t.Errorf("Error on value: %v, expected %v", got, want)
	}
}
