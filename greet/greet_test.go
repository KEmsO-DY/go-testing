package greet

import (
	"fmt"
	"testing"
)

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %v, expected %v", got, want)
	}
}
func assertError(t *testing.T, err error, errMsg string) {
	t.Helper()
	if err == nil {
		t.Error("Expected error")
	}
	if err.Error() != errMsg {
		t.Errorf("Got %v, expected %v", err.Error(), errMsg)
	}
}
func TestHello(t *testing.T) {
	type TestCase struct {
		input  string
		want   string
		errMsg *string
	}
	caseGroup := []TestCase{
		{"Go", "Hello, Go", nil},
		{" Go", "Hello,  Go", nil},
		{"World", "Hello, World", nil},
		{"Гофер", "Hello, Гофер", nil},
		{"", "", &ErrNameIsEmpty},
	}

	for _, c := range caseGroup {
		name := fmt.Sprintf("Testing %v", c.input)
		t.Run(name, func(t *testing.T) {
			got, err := Hello(c.input)
			assertString(t, got, c.want)
			if c.errMsg != nil {
				assertError(t, err, *c.errMsg)
			}
		})
	}
}
