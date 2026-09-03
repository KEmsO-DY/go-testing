package validate

import "errors"

var ErrEmptyName = errors.New("empty name")

func ValidateName(name string) error {
	if name == "" {
		return ErrEmptyName
	}
	return nil
}
