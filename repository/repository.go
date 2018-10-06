package repository

import (
	"github.com/dgraph-io/badger"
	"github.com/google/uuid"

	"github.com/mysza/paymentsapi/domain"
)

// PaymentsRepository provides access to the payments database.
type PaymentsRepository struct {
	db *badger.DB
}

// New creates a new repository using SQLite database.
func New(db *badger.DB) *PaymentsRepository {
	return &PaymentsRepository{db}
}

func (r *PaymentsRepository) set(payment *domain.Payment) error {
	encoded, err := domain.PaymentToByteSlice(payment)
	if err != nil {
		return err
	}
	return r.db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(payment.ID), encoded)
		return err
	})
}

// Add adds a payment to the database.
func (r *PaymentsRepository) Add(payment *domain.Payment) (string, error) {
	payment.ID = uuid.New().String()
	err := r.set(payment)
	if err != nil {
		return "", err
	}
	return payment.ID, nil
}

// Get retrieves single payment from the database.
func (r *PaymentsRepository) Get(id string) (*domain.Payment, error) {
	var encodedPayment []byte
	err := r.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(id))
		if err != nil {
			return err
		}
		encodedPayment, err = item.ValueCopy(nil)
		return err
	})
	if err != nil {
		return nil, err
	}
	return domain.PaymentFromByteSlice(encodedPayment)
}

// GetAll retrieves all payments from the database.
func (r *PaymentsRepository) GetAll() ([]*domain.Payment, error) {
	var encodedPayment []byte
	var payments []*domain.Payment
	err := r.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.IteratorOptions{
			PrefetchValues: true,
			PrefetchSize:   100,
		})
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			encodedPayment, err := it.Item().ValueCopy(encodedPayment)
			p, err := domain.PaymentFromByteSlice(encodedPayment)
			if err != nil {
				return err
			}
			payments = append(payments, p)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return payments, nil
}

// Update updates a payment in the database.
func (r *PaymentsRepository) Update(payment *domain.Payment) error {
	return r.set(payment)
}

// Delete deletes a payment from the database.
func (r *PaymentsRepository) Delete(id string) error {
	return r.db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(id))
	})
}

// Exists is a helper function to check if payment with give ID exists.
func (r *PaymentsRepository) Exists(id string) bool {
	payment, _ := r.Get(id)
	return payment != nil
}
