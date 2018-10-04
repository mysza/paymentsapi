package domain

import (
	"time"
)

// PaymentAttributes holds information about actual payment transaction
type PaymentAttributes struct {
	Amount               Amount             `json:"amount" validate:"required,numeric"`
	Beneficiary          BeneficiaryParty   `json:"beneficiary_party" validate:"required"`
	ChargesInformation   ChargesInformation `json:"charges_information" validate:"required"`
	Currency             Currency           `json:"currency" validate:"required,len=3,alpha"`
	Debtor               DebtorParty        `json:"debtor_party" validate:"required"`
	EndToEndReference    string             `json:"end_to_end_reference" validate:"required"`
	FX                   FX                 `json:"fx" validate:"required"`
	NumericReference     string             `json:"numeric_reference" validate:"required"`
	PaymentID            string             `json:"payment_id" validate:"required"`
	PaymentPurpose       string             `json:"payment_purpose" validate:"required"`
	PaymentScheme        string             `json:"payment_scheme" validate:"required"`
	PaymentType          string             `json:"payment_type" validate:"required"`
	ProcessingDate       time.Time          `json:"processing_date" validate:"required"`
	SchemePaymentType    string             `json:"scheme_payment_type" validate:"required"`
	SchemePaymentSubType string             `json:"scheme_payment_sub_type" validate:"required"`
	Reference            string             `json:"reference" validate:"required"`
	Sponsor              SponsorParty       `json:"sponsor_party" validate:"required"`
}
