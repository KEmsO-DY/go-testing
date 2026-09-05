package greet

import "errors"

var ErrNameIsEmpty string = "name cannot be empty"

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New(ErrNameIsEmpty)
	}
	return "Hello, " + name, nil
}
