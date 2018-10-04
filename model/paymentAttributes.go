package model

import (
	"time"
)

// PaymentAttributes holds information about actual payment transaction
type PaymentAttributes struct {
	Amount               Amount             `json:"amount"`
	Beneficiary          PaymentParty       `json:"beneficiary_party"`
	ChargesInformation   ChargesInformation `json:"charges_information"`
	Currency             Currency           `json:"currency"`
	Debtor               PaymentParty       `json:"debtor_party"`
	EndToEndReference    string             `json:"end_to_end_reference"`
	FX                   FX                 `json:"fx"`
	NumericReference     string             `json:"numeric_reference"`
	PaymentID            string             `json:"payment_id"`
	PaymentPurpose       string             `json:"payment_purpose"`
	PaymentScheme        string             `json:"payment_scheme"`
	PaymentType          string             `json:"payment_type"`
	ProcessingDate       time.Time          `json:"processing_date"`
	SchemePaymentType    string             `json:"scheme_payment_type"`
	SchemePaymentSubType string             `json:"scheme_payment_sub_type"`
	Reference            string             `json:"reference"`
}
