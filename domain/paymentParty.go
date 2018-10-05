package domain

import validator "gopkg.in/go-playground/validator.v9"

// Account holds the base information about an account
type Account struct {
	AccountNumber string `json:"account_number" validate:"required"`
	BankID        string `json:"bank_id" validate:"required"`
	BankIDCode    string `json:"bank_id_code" validate:"required"`
}

// Validate validates if a given Account object is valid.
func (acc Account) Validate(v *validator.Validate) error {
	return v.Struct(acc)
}

// PaymentParty is a party taking part in the payment transaction
type PaymentParty struct {
	Account
	AccountName       string `json:"account_name" validate:"required"`
	AccountNumberCode string `json:"account_number_code" validate:"required,oneof=IBAN BBAN"`
	Address           string `json:"address" validate:"required"`
	Name              string `json:"name" validate:"required"`
}

// Validate validates if a given PaymentParty object is valid.
func (pp PaymentParty) Validate(v *validator.Validate) error {
	return v.Struct(pp)
}

// BeneficiaryPaymentParty is a PaymentParty with additional AccountType field
type BeneficiaryPaymentParty struct {
	PaymentParty
	AccountType int `json:"account_type,omitempty" validate:"isdefault"`
}

// Validate validates if a given BeneficiaryPaymentParty object is valid.
func (bpp BeneficiaryPaymentParty) Validate(v *validator.Validate) error {
	return v.Struct(bpp)
}
