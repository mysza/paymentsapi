package domain

import validator "gopkg.in/go-playground/validator.v9"

// Validatable is an interface to be implemented by the structures
// that are self-validatable.
type Validatable interface {
	Validate(v *validator.Validate) error
}
