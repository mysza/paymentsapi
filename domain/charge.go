package domain

// Charge represents payment charge
type Charge struct {
	// Amount is the charged amount; is required and must be a number
	Amount Amount `json:"amount" validate:"required,numeric"`
	// Currency is the currency the amount was charged with, ISO 4217 3-letter string
	Currency Currency `json:"currency" validate:"required,len=3,alpha"`
}
