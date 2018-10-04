package domain

import (
	"testing"

	validator "gopkg.in/go-playground/validator.v9"
)

func TestSponsorPartyFieldsValidationRules(t *testing.T) {
	var scenarios = []struct {
		description string
		party       SponsorParty
		passed      func(error) bool
	}{
		{
			description: "No error if all fields correct",
			party: SponsorParty{
				AccountNumber: "56781234",
				BankID:        "123123",
				BankIDCode:    "GBDSC",
			},
			passed: func(err error) bool { return err == nil },
		},
		{
			description: "Error if AccountNumber missing",
			party: SponsorParty{
				BankID:     "123123",
				BankIDCode: "GBDSC",
			},
			passed: func(err error) bool { return err != nil },
		},
		{
			description: "Error if BankID missing",
			party: SponsorParty{
				AccountNumber: "56781234",
				BankIDCode:    "GBDSC",
			},
			passed: func(err error) bool { return err != nil },
		},
		{
			description: "Error if BankIDCode missing",
			party: SponsorParty{
				AccountNumber: "56781234",
				BankID:        "123123",
			},
			passed: func(err error) bool { return err != nil },
		},
	}
	validator := validator.New()
	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			if err := validator.Struct(scenario.party); !scenario.passed(err) {
				t.Errorf("Validation is defined inproperly: %s", err)
			}
		})
	}
}
