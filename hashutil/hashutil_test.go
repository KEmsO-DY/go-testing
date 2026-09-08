package hashutil

import (
	"fmt"
	"testing"
)

func TestHashSHA256(t *testing.T) {
	type Case struct {
		s    string
		want string
	}
	TestCases := []Case{
		{"", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		{"hello", "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"},
		{"Привет", "dd679c0b9fd408a04148aa7d30c9df393f67b7227f65693fffe0ed6d0f0ade59"},
	}

	for _, c := range TestCases {
		c := c
		name := fmt.Sprintf("Testing_input:%v", c.s)
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := HashSHA256(c.s); got != c.want {
				t.Errorf("Got:%v, want:%v", got, c.want)
			}
		})
	}
}
