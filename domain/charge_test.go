package domain

import (
	"testing"

	validator "gopkg.in/go-playground/validator.v9"
)

func TestRequiredFields(t *testing.T) {
	var scenarios = []struct {
		description string
		charge      Charge
		passed      func(error) bool
	}{
		{
			description: "No error if both fields correct",
			charge:      Charge{Currency: "USD", Amount: "123.22"},
			passed:      func(err error) bool { return err == nil },
		},
		{
			description: "Error if currency empty",
			charge:      Charge{Amount: "123.22"},
			passed:      func(err error) bool { return err != nil },
		},
		{
			description: "Error if Amount empty",
			charge:      Charge{Currency: "USD"},
			passed:      func(err error) bool { return err != nil },
		},
		{
			description: "Error if Amount non-numeric",
			charge:      Charge{Currency: "USD", Amount: "abc"},
			passed:      func(err error) bool { return err != nil },
		},
		{
			description: "Error if Currency non-alpha",
			charge:      Charge{Currency: "US3", Amount: "abc"},
			passed:      func(err error) bool { return err != nil },
		},
		{
			description: "Error if Currency length != 3",
			charge:      Charge{Currency: "US", Amount: "abc"},
			passed:      func(err error) bool { return err != nil },
		},
	}
	validator := validator.New()
	t.Run("Requires the amount to be set", func(t *testing.T) {
		for _, scenario := range scenarios {
			t.Run(scenario.description, func(t *testing.T) {
				if err := validator.Struct(scenario.charge); !scenario.passed(err) {
					t.Errorf("Validation is defined inproperly: %s", err)
				}
			})
		}
	})
}
