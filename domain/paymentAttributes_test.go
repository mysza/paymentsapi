package domain

import (
	"testing"
	"time"

	validator "gopkg.in/go-playground/validator.v9"
)

func TestPaymentAttributesFieldsValidationRules(t *testing.T) {
	var scenarios = []struct {
		description string
		attributes  PaymentAttributes
		passed      func(error) bool
	}{
		{
			description: "No error if all fields correct",
			attributes: PaymentAttributes{
				Amount:               "100.12",
				Beneficiary:          BeneficiaryPaymentParty{},
				ChargesInformation:   ChargesInformation{},
				Currency:             "USD",
				Debtor:               PaymentParty{},
				EndToEndReference:    "Some generic string",
				FX:                   FX{},
				NumericReference:     "123456",
				PaymentID:            "123456789012345678",
				PaymentPurpose:       "Paying for goods/services",
				PaymentScheme:        "FPS",
				PaymentType:          "Credit",
				ProcessingDate:       time.Now(),
				SchemePaymentType:    "InternetBanking",
				SchemePaymentSubType: "ImmediatePayment",
				Reference:            "Payment for Em's piano lessons",
				Sponsor:              Account{},
			},
			passed: func(err error) bool { return err == nil },
		},
	}
	validator := validator.New()
	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			if err := validator.Struct(scenario.attributes); !scenario.passed(err) {
				t.Errorf("Validation is defined inproperly: %s", err)
			}
		})
	}
}
