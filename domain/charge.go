package domain

// Charge represents payment charge
type Charge struct {
	Amount   string `json:"amount" validate:"required,numeric"`       // Amount is the charged amount; is required and must be a number
	Currency string `json:"currency" validate:"required,len=3,alpha"` // Currency is the currency the amount was charged with, ISO 4217 3-letter string
}
