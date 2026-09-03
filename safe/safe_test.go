package safe

import "testing"

func TestMustAt(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Recover is nil")
		} else if r != "index out of range" {
			t.Errorf("Unsuspected name of panic")
		}
	}()
	slice := make([]int, 1)
	_ = MustAt(slice, -5)
}
