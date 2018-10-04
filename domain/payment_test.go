package domain

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	validator "gopkg.in/go-playground/validator.v9"
)

func TestPaymentFieldsValidationRules(t *testing.T) {
	var scenarios = []struct {
		description string
		payment     Payment
		passed      func(error) bool
	}{
		{
			description: "No error if all fields correct",
			payment: Payment{
				ID:             uuid.New(),
				OrganisationID: uuid.New(),
				Attributes:     PaymentAttributes{},
			},
			passed: func(err error) bool { return err == nil },
		},
	}
	validator := validator.New()
	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			fmt.Printf("Payment: %#v", scenario.payment)
			if err := validator.Struct(scenario.payment); !scenario.passed(err) {
				t.Errorf("Validation is defined inproperly: %s", err)
			}
		})
	}
}
