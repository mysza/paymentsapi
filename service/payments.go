package service

import (
	"fmt"

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
	Exists(string) bool
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

func (ps *PaymentsService) validate(payment *domain.Payment) error {
	return ps.validator.Struct(payment)
}

// Add adds a new payment to the service.
// Before that, it validates the argument.
func (ps *PaymentsService) Add(payment *domain.Payment) (string, error) {
	// check if ID is set
	if payment.ID != "" {
		return "", NewInputError("Payment cannot have ID set when adding to repository")
	}
	if err := ps.validate(payment); err != nil {
		return "", NewInputError(err.Error())
	}
	return ps.repo.Add(payment)
}

// GetAll simply returns all payments from the repository.
func (ps *PaymentsService) GetAll() ([]*domain.Payment, error) {
	return ps.repo.GetAll()
}

// Update updates existing payment.
func (ps *PaymentsService) Update(payment *domain.Payment) error {
	if err := ps.validate(payment); err != nil {
		return NewInputError(err.Error())
	}
	if exists := ps.repo.Exists(payment.ID); !exists {
		return NewNotFoundError(fmt.Sprintf("Payment with ID: %v does not exist", payment.ID))
	}
	return ps.repo.Update(payment)
}

// Get retrieves a single Payment based on ID
func (ps *PaymentsService) Get(id string) (*domain.Payment, error) {
	if id == "" {
		return nil, NewInputError("Invalid ID")
	}
	payment, err := ps.repo.Get(id)
	if err != nil {
		return nil, NewNotFoundError(fmt.Sprintf("Payment with id %v not found", id))
	}
	return payment, nil
}

// Delete deletes payment with given ID from the repository.
func (ps *PaymentsService) Delete(id string) error {
	if id == "" {
		return NewInputError("Invalid ID")
	}
	if exists := ps.repo.Exists(id); !exists {
		return NewNotFoundError(fmt.Sprintf("Payment with ID %v does not exist", id))
	}
	return ps.repo.Delete(id)
}
