package model

// PaymentParty describes a single party of the payment transaction
type PaymentParty struct {
	AccountNumber     string            `json:"account_number"`
	BankID            string            `json:"bank_id"`
	BankIDCode        string            `json:"bank_id_code"`
	AccountName       string            `json:"account_name,omitempty"`
	AccountNumberCode AccountNumberCode `json:"account_number_code,omitempty"`
	Address           string            `json:"address,omitempty"`
	Name              string            `json:"name,omitempty"`
	AccountType       int               `json:"account_type,omitempty"`
}
