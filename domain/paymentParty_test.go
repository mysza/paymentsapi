package domain

import (
	"testing"

	validator "gopkg.in/go-playground/validator.v9"
)

func TestAccountFieldsValidationRules(t *testing.T) {
	var scenarios = []struct {
		description string
		account     *Account
		passed      func(error) bool
	}{
		{
			description: "No error if all fields correct",
			account: &Account{
				AccountNumber: "56781234",
				BankID:        "123123",
				BankIDCode:    "GBDSC",
			},
			passed: func(err error) bool { return err == nil },
		},
		{
			description: "Error if AccountNumber missing",
			account: &Account{
				BankID:     "123123",
				BankIDCode: "GBDSC",
			},
			passed: func(err error) bool { return err != nil },
		},
		{
			description: "Error if BankID missing",
			account: &Account{
				AccountNumber: "56781234",
				BankIDCode:    "GBDSC",
			},
			passed: func(err error) bool { return err != nil },
		},
		{
			description: "Error if BankIDCode missing",
			account: &Account{
				AccountNumber: "56781234",
				BankID:        "123123",
			},
			passed: func(err error) bool { return err != nil },
		},
	}
	validator := validator.New()
	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			if err := scenario.account.Validate(validator); !scenario.passed(err) {
				t.Errorf("Validation is defined inproperly: %s", err)
			}
		})
	}
}

func TestPaymentPartyFieldsValidationRules(t *testing.T) {
	var scenarios = []struct {
		description string
		party       *PaymentParty
		passed      func(error) bool
	}{
		{
			description: "No error if all fields correct",
			party: &PaymentParty{
				Account: &Account{
					AccountNumber: "56781234",
					BankID:        "123123",
					BankIDCode:    "GBDSC",
				},
				AccountName:       "EJ Brown Black",
				AccountNumberCode: "IBAN",
				Address:           "10 Debtor Crescent Sourcetown NE1",
				Name:              "EJ Brown Black",
			},
			passed: func(err error) bool { return err == nil },
		},
		{
			description: "Error if AccountName missing",
			party: &PaymentParty{
				Account: &Account{
					AccountNumber: "56781234",
					BankID:        "123123",
					BankIDCode:    "GBDSC",
				},
				AccountNumberCode: "IBAN",
				Address:           "10 Debtor Crescent Sourcetown NE1",
				Name:              "EJ Brown Black",
			},
			passed: func(err error) bool { return err != nil },
		},
		{
			description: "Error if AccountNumberCode missing",
			party: &PaymentParty{
				Account: &Account{
					AccountNumber: "56781234",
					BankID:        "123123",
					BankIDCode:    "GBDSC",
				},
				AccountName: "EJ Brown Black",
				Address:     "10 Debtor Crescent Sourcetown NE1",
				Name:        "EJ Brown Black",
			},
			passed: func(err error) bool { return err != nil },
		},
		{
			description: "Error if AccountNumberCode other than from oneof set",
			party: &PaymentParty{
				Account: &Account{
					AccountNumber: "56781234",
					BankID:        "123123",
					BankIDCode:    "GBDSC",
				},
				AccountNumberCode: "ICCA",
				AccountName:       "EJ Brown Black",
				Address:           "10 Debtor Crescent Sourcetown NE1",
				Name:              "EJ Brown Black",
			},
			passed: func(err error) bool { return err != nil },
		},
		{
			description: "Error if Address missing",
			party: &PaymentParty{
				Account: &Account{
					AccountNumber: "56781234",
					BankID:        "123123",
					BankIDCode:    "GBDSC",
				},
				AccountNumberCode: "IBAN",
				AccountName:       "EJ Brown Black",
				Name:              "EJ Brown Black",
			},
			passed: func(err error) bool { return err != nil },
		},
		{
			description: "Error if Name missing",
			party: &PaymentParty{
				Account: &Account{
					AccountNumber: "56781234",
					BankID:        "123123",
					BankIDCode:    "GBDSC",
				},
				AccountNumberCode: "IBAN",
				AccountName:       "EJ Brown Black",
				Address:           "10 Debtor Crescent Sourcetown NE1",
			},
			passed: func(err error) bool { return err != nil },
		},
	}
	validator := validator.New()
	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			if err := scenario.party.Validate(validator); !scenario.passed(err) {
				t.Errorf("Validation is defined inproperly: %s", err)
			}
		})
	}
}

func TestBeneficiaryPaymentPartyFieldsValidationRules(t *testing.T) {
	var scenarios = []struct {
		description string
		party       *BeneficiaryPaymentParty
		passed      func(error) bool
	}{
		{
			description: "No error if all fields correct",
			party: &BeneficiaryPaymentParty{
				PaymentParty: &PaymentParty{
					Account: &Account{
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
			passed: func(err error) bool { return err == nil },
		},
		{
			description: "Error if AccountType != 0",
			party: &BeneficiaryPaymentParty{
				PaymentParty: &PaymentParty{
					Account: &Account{
						AccountNumber: "56781234",
						BankID:        "123123",
						BankIDCode:    "GBDSC",
					},
					AccountName:       "EJ Brown Black",
					AccountNumberCode: "IBAN",
					Address:           "10 Debtor Crescent Sourcetown NE1",
					Name:              "EJ Brown Black",
				},
				AccountType: 10,
			},
			passed: func(err error) bool { return err != nil },
		},
	}
	validator := validator.New()
	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			if err := scenario.party.Validate(validator); !scenario.passed(err) {
				t.Errorf("Validation is defined inproperly: %s", err)
			}
		})
	}
}
