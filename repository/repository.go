package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/mysza/paymentsapi/domain"
)

// PaymentsRepository provides access to the payments database.
type PaymentsRepository struct {
	db *gorm.DB
}

// New creates a new repository using SQLite database.
func New(db *gorm.DB) *PaymentsRepository {
	payment := &domain.Payment{}
	attributes := &domain.PaymentAttributes{}
	beneficiary := &domain.BeneficiaryPaymentParty{}
	charges := &domain.ChargesInformation{}
	debtor := &domain.PaymentParty{}
	fx := &domain.FX{}
	sponsor := &domain.Account{}
	charge := &domain.Charge{}
	db.AutoMigrate(payment, attributes, beneficiary, charges, debtor, fx, sponsor, charge)
	return &PaymentsRepository{db}
}

// Add adds a payment to the database.
func (r *PaymentsRepository) Add(payment *domain.Payment) (string, error) {
	err := r.db.Create(payment).Error
	if err != nil {
		return "", err
	}
	return payment.ID, nil
}

// GetAll retrieves all payments from the database.
func (r *PaymentsRepository) GetAll() ([]*domain.Payment, error) {
	var payments []*domain.Payment
	err := r.db.Find(&payments).Error
	if err != nil {
		return nil, err
	}
	return payments, nil
}

// Update updates a payment in the database.
func (r *PaymentsRepository) Update(payment *domain.Payment) error {
	return r.db.Save(payment).Error
}

// Get retrieves single payment from the database.
func (r *PaymentsRepository) Get(id string) (*domain.Payment, error) {
	var payment *domain.Payment
	err := r.db.Where("ID = ?", id).First(payment).Error
	if err != nil {
		return nil, err
	}
	return payment, nil
}

// Delete deletes a payment from the database.
func (r *PaymentsRepository) Delete(id string) error {
	return r.db.Where("ID = ?", id).Delete(domain.Payment{}).Error
}

// Exists checks whether a payment exists in the database.
func (r *PaymentsRepository) Exists(id string) bool {
	var payment *domain.Payment
	r.db.Where("ID = ?", id).First(payment)
	return payment != nil
}
