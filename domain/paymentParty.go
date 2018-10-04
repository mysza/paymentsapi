package domain

type partyBase struct {
	AccountNumber string `json:"account_number" validate:"required"`
	BankID        string `json:"bank_id" validate:"required"`
	BankIDCode    string `json:"bank_id_code" validate:"required"`
}

type paymentParty struct {
	partyBase
	AccountName       string            `json:"account_name" validate:"required"`
	AccountNumberCode AccountNumberCode `json:"account_number_code" validate:"required,oneof=IBAN BBAN"`
	Address           string            `json:"address" validate:"required"`
	Name              string            `json:"name" validate:"required"`
}

type paymentPartyExtended struct {
	paymentParty
	AccountType int `json:"account_type,omitempty" validate:"default"`
}

// SponsorParty describes the sponsor of a payment
type SponsorParty partyBase

// DebtorParty describes the debtor of a payment
type DebtorParty paymentParty

// BeneficiaryParty describes the beneficiary of a payment
type BeneficiaryParty paymentPartyExtended
