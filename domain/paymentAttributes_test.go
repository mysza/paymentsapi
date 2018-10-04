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
				Beneficiary: BeneficiaryPaymentParty{
					PaymentParty: PaymentParty{
						Account: Account{
							AccountNumber: "56781234",
							BankID:        "123123",
							BankIDCode:    "GBDSC",
						},
						AccountName:       "EJ Brown Black",
						AccountNumberCode: "IBAN",
						Address:           "10 Debtor Crescent Sourcetown NE1",
						Name:              "EJ Brown Black",
					},
					AccountType: 0,
				},
				Debtor: PaymentParty{
					Account: Account{
						AccountNumber: "56781234",
						BankID:        "123123",
						BankIDCode:    "GBDSC",
					},
					AccountName:       "EJ Brown Black",
					AccountNumberCode: "IBAN",
					Address:           "10 Debtor Crescent Sourcetown NE1",
					Name:              "EJ Brown Black",
				},
				Sponsor: Account{
					AccountNumber: "56781234",
					BankID:        "123123",
					BankIDCode:    "GBDSC",
				},
				ChargesInformation: ChargesInformation{
					BearerCode:              "SHAR",
					ReceiverChargesAmount:   "100.12",
					ReceiverChargesCurrency: "USD",
					SenderCharges: []Charge{
						Charge{Currency: "USD", Amount: "5.00"},
						Charge{Currency: "GBP", Amount: "15.00"},
					},
				},
				FX: FX{
					ContractReference: "FX123",
					ExchangeRate:      "2.00",
					OriginalAmount:    "100.12",
					OriginalCurrency:  "USD",
				},
				ProcessingDate:       time.Now(),
				Amount:               "100.12",
				Currency:             "USD",
				EndToEndReference:    "Some generic string",
				NumericReference:     "123456",
				PaymentID:            "123456789012345678",
				PaymentPurpose:       "Paying for goods/services",
				PaymentScheme:        "FPS",
				PaymentType:          "Credit",
				SchemePaymentType:    "InternetBanking",
				SchemePaymentSubType: "ImmediatePayment",
				Reference:            "Payment for Em's piano lessons",
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
