package domain

// Account holds the base information about an account
type Account struct {
	AccountNumber string `json:"account_number" validate:"required"`
	BankID        string `json:"bank_id" validate:"required"`
	BankIDCode    string `json:"bank_id_code" validate:"required"`
}

// PaymentParty is a party taking part in the payment transaction
type PaymentParty struct {
	Account
	AccountName       string `json:"account_name" validate:"required"`
	AccountNumberCode string `json:"account_number_code" validate:"required,oneof=IBAN BBAN"`
	Address           string `json:"address" validate:"required"`
	Name              string `json:"name" validate:"required"`
}

// BeneficiaryPaymentParty is a PaymentParty with additional AccountType field
type BeneficiaryPaymentParty struct {
	PaymentParty
	AccountType int `json:"account_type,omitempty" validate:"isdefault"`
}
