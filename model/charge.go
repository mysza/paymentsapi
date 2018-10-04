package model

// Charge represents payment charge
type Charge struct {
	Amount   Amount   `json:"amount"`   // Amount is the charged amount
	Currency Currency `json:"currency"` // Currency is the currency the amount was charged with
}
