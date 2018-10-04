package domain

import (
	"testing"

	validator "gopkg.in/go-playground/validator.v9"
)

func TestFieldValidationRules(t *testing.T) {
	var scenarios = []struct {
		description string
		chargesInfo ChargesInformation
		passed      func(error) bool
	}{
		{
			description: "No error if all fields correct",
			chargesInfo: ChargesInformation{
				BearerCode:              "SHAR",
				ReceiverChargesAmount:   "100.12",
				ReceiverChargesCurrency: "USD",
				SenderCharges: []Charge{
					Charge{Currency: "USD", Amount: "5.00"},
					Charge{Currency: "GBP", Amount: "15.00"},
				},
			},
			passed: func(err error) bool { return err == nil },
		},
		{
			description: "Error if BearerCode empty",
			chargesInfo: ChargesInformation{
				BearerCode:              "",
				ReceiverChargesAmount:   "100.12",
				ReceiverChargesCurrency: "USD",
				SenderCharges: []Charge{
					Charge{Currency: "USD", Amount: "5.00"},
					Charge{Currency: "GBP", Amount: "15.00"},
				},
			},
			passed: func(err error) bool { return err != nil },
		},
		{
			description: "Error if Amount empty",
			chargesInfo: ChargesInformation{
				BearerCode:              "SHAR",
				ReceiverChargesAmount:   "",
				ReceiverChargesCurrency: "USD",
				SenderCharges: []Charge{
					Charge{Currency: "USD", Amount: "5.00"},
					Charge{Currency: "GBP", Amount: "15.00"},
				},
			},
			passed: func(err error) bool { return err != nil },
		},
		{
			description: "Error if Amount non-numeric",
			chargesInfo: ChargesInformation{
				BearerCode:              "SHAR",
				ReceiverChargesAmount:   "abc",
				ReceiverChargesCurrency: "USD",
				SenderCharges: []Charge{
					Charge{Currency: "USD", Amount: "5.00"},
					Charge{Currency: "GBP", Amount: "15.00"},
				},
			},
			passed: func(err error) bool { return err != nil },
		},
		{
			description: "Error if Currency empty",
			chargesInfo: ChargesInformation{
				BearerCode:              "SHAR",
				ReceiverChargesAmount:   "100.12",
				ReceiverChargesCurrency: "",
				SenderCharges: []Charge{
					Charge{Currency: "USD", Amount: "5.00"},
					Charge{Currency: "GBP", Amount: "15.00"},
				},
			},
			passed: func(err error) bool { return err != nil },
		},
		{
			description: "Error if Currency length != 3",
			chargesInfo: ChargesInformation{
				BearerCode:              "SHAR",
				ReceiverChargesAmount:   "100.12",
				ReceiverChargesCurrency: "US",
				SenderCharges: []Charge{
					Charge{Currency: "USD", Amount: "5.00"},
					Charge{Currency: "GBP", Amount: "15.00"},
				},
			},
			passed: func(err error) bool { return err != nil },
		},
		{
			description: "Error if Currency non-alpha",
			chargesInfo: ChargesInformation{
				BearerCode:              "SHAR",
				ReceiverChargesAmount:   "100.12",
				ReceiverChargesCurrency: "US3",
				SenderCharges: []Charge{
					Charge{Currency: "USD", Amount: "5.00"},
					Charge{Currency: "GBP", Amount: "15.00"},
				},
			},
			passed: func(err error) bool { return err != nil },
		},
		{
			description: "No error if SenderCharges empty",
			chargesInfo: ChargesInformation{
				BearerCode:              "SHAR",
				ReceiverChargesAmount:   "100.12",
				ReceiverChargesCurrency: "USD",
				SenderCharges:           []Charge{},
			},
			passed: func(err error) bool { return err == nil },
		},
		{
			description: "Error if SenderCharges Charge invalid",
			chargesInfo: ChargesInformation{
				BearerCode:              "SHAR",
				ReceiverChargesAmount:   "100.12",
				ReceiverChargesCurrency: "US3",
				SenderCharges:           []Charge{Charge{}},
			},
			passed: func(err error) bool { return err != nil },
		},
	}
	validator := validator.New()
	t.Run("Requires the amount to be set", func(t *testing.T) {
		for _, scenario := range scenarios {
			t.Run(scenario.description, func(t *testing.T) {
				if err := validator.Struct(scenario.chargesInfo); !scenario.passed(err) {
					t.Errorf("Validation is defined inproperly: %s", err)
				}
			})
		}
	})
}
