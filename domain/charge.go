package domain

import validator "gopkg.in/go-playground/validator.v9"

// Charge represents payment charge
type Charge struct {
	// Amount is the charged amount; is required and must be a number
	Amount string `json:"amount" validate:"required,numeric"`
	// Currency is the currency the amount was charged with, ISO 4217 3-letter string
	Currency string `json:"currency" validate:"required,len=3,alpha"`
}

// Validate validates if a given Charge object is valid.
func (c Charge) Validate(v *validator.Validate) error {
	return v.Struct(c)
}
