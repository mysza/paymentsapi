package domain

import (
	"time"

	validator "gopkg.in/go-playground/validator.v9"
)

// PaymentAttributes holds information about actual payment transaction
type PaymentAttributes struct {
	Amount               string                  `json:"amount" validate:"required,numeric"`
	Beneficiary          BeneficiaryPaymentParty `json:"beneficiary_party" validate:"required"`
	ChargesInformation   ChargesInformation      `json:"charges_information" validate:"required"`
	Currency             string                  `json:"currency" validate:"required,len=3,alpha"`
	Debtor               PaymentParty            `json:"debtor_party" validate:"required"`
	EndToEndReference    string                  `json:"end_to_end_reference" validate:"required"`
	FX                   FX                      `json:"fx" validate:"required"`
	NumericReference     string                  `json:"numeric_reference" validate:"required"`
	PaymentID            string                  `json:"payment_id" validate:"required"`
	PaymentPurpose       string                  `json:"payment_purpose" validate:"required"`
	PaymentScheme        string                  `json:"payment_scheme" validate:"required"`
	PaymentType          string                  `json:"payment_type" validate:"required"`
	ProcessingDate       time.Time               `json:"processing_date" validate:"required"`
	SchemePaymentType    string                  `json:"scheme_payment_type" validate:"required"`
	SchemePaymentSubType string                  `json:"scheme_payment_sub_type" validate:"required"`
	Reference            string                  `json:"reference" validate:"required"`
	Sponsor              Account                 `json:"sponsor_party" validate:"required"`
}

// Validate validates if a given PaymentAttributes object is valid.
func (pa PaymentAttributes) Validate(v *validator.Validate) error {
	return v.Struct(pa)
}
