package model

// AccountNumberCode is IBAN or BBAN
type AccountNumberCode string

const (
	// IBAN account number code
	IBAN AccountNumberCode = "IBAN"
	// BBAN account number code
	BBAN AccountNumberCode = "BBAN"
)
