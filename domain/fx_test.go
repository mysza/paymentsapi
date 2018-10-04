package domain

import (
	"testing"

	validator "gopkg.in/go-playground/validator.v9"
)

func TestFXFieldsValidationRules(t *testing.T) {
	var scenarios = []struct {
		description string
		fx          FX
		passed      func(error) bool
	}{
		{
			description: "No error if all fields correct",
			fx: FX{
				ContractReference: "FX123",
				ExchangeRate:      "2.00",
				OriginalAmount:    "100.12",
				OriginalCurrency:  "USD",
			},
			passed: func(err error) bool { return err == nil },
		},
		{
			description: "Error if ContractReference empty",
			fx: FX{
				ExchangeRate:     "2.00",
				OriginalAmount:   "100.12",
				OriginalCurrency: "USD",
			},
			passed: func(err error) bool { return err != nil },
		},
		{
			description: "Error if ContractReference non-alphanumeric",
			fx: FX{
				ContractReference: "FX%$*",
				ExchangeRate:      "2.00",
				OriginalAmount:    "100.12",
				OriginalCurrency:  "USD",
			},
			passed: func(err error) bool { return err != nil },
		},
		{
			description: "Error if ExchangeRate empty",
			fx: FX{
				ContractReference: "FX123",
				OriginalAmount:    "100.12",
				OriginalCurrency:  "USD",
			},
			passed: func(err error) bool { return err != nil },
		},
		{
			description: "Error if ExchangeRate non-numeric",
			fx: FX{
				ContractReference: "FX123",
				ExchangeRate:      "abc",
				OriginalAmount:    "100.12",
				OriginalCurrency:  "USD",
			},
			passed: func(err error) bool { return err != nil },
		},
		{
			description: "Error if OriginalAmount empty",
			fx: FX{
				ContractReference: "FX123",
				ExchangeRate:      "2.00",
				OriginalCurrency:  "USD",
			},
			passed: func(err error) bool { return err != nil },
		},
		{
			description: "Error if OriginalAmount non-numeric",
			fx: FX{
				ContractReference: "FX123",
				ExchangeRate:      "2.00",
				OriginalAmount:    "100.aa",
				OriginalCurrency:  "USD",
			},
			passed: func(err error) bool { return err != nil },
		},
		{
			description: "Error if OriginalCurrency empty",
			fx: FX{
				ContractReference: "FX123",
				ExchangeRate:      "2.00",
				OriginalAmount:    "100.12",
			},
			passed: func(err error) bool { return err != nil },
		},
		{
			description: "Error if OriginalCurrency length != 3",
			fx: FX{
				ContractReference: "FX123",
				ExchangeRate:      "2.00",
				OriginalAmount:    "100.12",
				OriginalCurrency:  "US",
			},
			passed: func(err error) bool { return err != nil },
		},
		{
			description: "Error if OriginalCurrency non-alpha",
			fx: FX{
				ContractReference: "FX123",
				ExchangeRate:      "2.00",
				OriginalAmount:    "100.12",
				OriginalCurrency:  "US3",
			},
			passed: func(err error) bool { return err != nil },
		},
	}
	validator := validator.New()
	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			if err := validator.Struct(scenario.fx); !scenario.passed(err) {
				t.Errorf("Validation is defined inproperly: %s", err)
			}
		})
	}
}
