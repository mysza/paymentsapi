package service

import (
	"testing"

	"github.com/google/uuid"
	"github.com/mysza/paymentsapi/domain"
)

func TestAddReturnsErrorOnInvalidInput(t *testing.T) {
	t.Run("PaymentsService Add returns error if invalid input passed", func(t *testing.T) {
		ps := NewPaymentsService(nil)
		payment := &domain.Payment{
			ID:             uuid.New(),
			OrganisationID: uuid.New(),
		}
		if _, err := ps.Add(payment); err == nil {
			t.Error("Payments service didn't return error despite invalid input")
		}
	})
}
