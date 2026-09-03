package validate

import "testing"

func TestValidateName(t *testing.T) {
	Arg := ""
	if err := ValidateName(Arg); err == nil {
		t.Errorf("Error is nil")
	} else if err != ErrEmptyName {
		t.Errorf("Unsuspected error")
	}
}
