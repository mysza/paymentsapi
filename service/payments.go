package service

import (
	"github.com/mysza/paymentsapi/domain"
	validator "gopkg.in/go-playground/validator.v9"
)

// PaymentsRepository is an interface that any repository
// that should be used by the service for Payments storage
// need to implement.
type PaymentsRepository interface {
	Add(*domain.Payment) (string, error)
	GetAll() ([]*domain.Payment, error)
	Update(*domain.Payment) error
	Get(string) (*domain.Payment, error)
	Delete(string) error
}

// PaymentsService implements all use cases of the Payments API.
type PaymentsService struct {
	repo      PaymentsRepository
	validator *validator.Validate
}

// NewPaymentsService creates a new instance of PaymentsService
// with the provided repository.
func NewPaymentsService(repo PaymentsRepository) *PaymentsService {
	return &PaymentsService{
		repo:      repo,
		validator: validator.New(),
	}
}

// Add adds a new payment to the service.
func (ps *PaymentsService) Add(payment *domain.Payment) (string, error) {
	if err := payment.Validate(ps.validator); err != nil {
		return "", err
	}
	return ps.repo.Add(payment)
}
