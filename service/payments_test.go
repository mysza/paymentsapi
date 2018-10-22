package service

import (
	"path/filepath"
	"testing"

	"github.com/jinzhu/copier"
	"github.com/mysza/paymentsapi/domain"
	"github.com/mysza/paymentsapi/service/mocks"
	"github.com/mysza/paymentsapi/test"
	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	assert := assert.New(t)
	testDataDir := filepath.Join("..", "testdata")
	validPayment := test.PaymentFromFile(t, filepath.Join(testDataDir, "validPayment.json"))
	invalidPayment := test.PaymentFromFile(t, filepath.Join(testDataDir, "invalidPayment.json"))

	t.Run("Add payment", func(t *testing.T) {
		t.Run("PaymentsService Add returns error if invalid input passed", func(t *testing.T) {
			ps := NewPaymentsService(nil)

			_, err := ps.Add(invalidPayment)

			assert.Error(err)
		})

		t.Run("PaymentsService Add returns added payment ID if the input was valid", func(t *testing.T) {
			var payment = domain.Payment{}
			copier.Copy(&payment, validPayment)
			payment.ID = ""
			repo := new(mocks.PaymentsRepository)
			repo.On("Add", &payment).Return("4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43", nil)
			ps := NewPaymentsService(repo)

			id, err := ps.Add(&payment)

			assert.Equal("4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43", id, "Returned ID should equal expected value")
			assert.Nil(err)
			repo.AssertExpectations(t)
		})

		t.Run("PaymentsService Add returns error if payment ID was set", func(t *testing.T) {
			ps := NewPaymentsService(nil)

			id, err := ps.Add(validPayment)

			assert.Empty(id, "Returned ID should be empty")
			assert.Error(err, "Error should be set")
		})
	})

	t.Run("Get all payments", func(t *testing.T) {
		payments := []*domain.Payment{validPayment, validPayment}
		repo := new(mocks.PaymentsRepository)
		repo.On("GetAll").Return(payments, nil)
		ps := NewPaymentsService(repo)

		retPayments, err := ps.GetAll()

		assert.Equal(payments, retPayments, "Returned payments should be all in repository")
		assert.Nil(err)
		repo.AssertExpectations(t)
	})

	t.Run("Update payment", func(t *testing.T) {

		t.Run("PaymentsService Update returns error if invalid input passed", func(t *testing.T) {
			ps := NewPaymentsService(nil)

			err := ps.Update(invalidPayment)

			assert.Error(err)
		})

		t.Run("PaymentsService Update returns error if invalid input passed", func(t *testing.T) {
			repo := new(mocks.PaymentsRepository)
			repo.On("Exists", validPayment.ID).Return(false)
			ps := NewPaymentsService(repo)

			err := ps.Update(validPayment)

			assert.Error(err, "Update with non-existing payment shoud return error")
			repo.AssertExpectations(t)
		})

		t.Run("PaymentsService Update returns nil if the input was valid", func(t *testing.T) {
			repo := new(mocks.PaymentsRepository)
			repo.On("Exists", validPayment.ID).Return(true)
			repo.On("Update", validPayment).Return(nil)
			ps := NewPaymentsService(repo)

			err := ps.Update(validPayment)

			assert.Nil(err)
			repo.AssertExpectations(t)
		})
	})

	t.Run("Get payment", func(t *testing.T) {
		ps := NewPaymentsService(nil)

		t.Run("PaymentsService Get returns error if invalid ID passed", func(t *testing.T) {
			retPayment, err := ps.Get("")

			assert.Error(err)
			assert.Empty(retPayment)
		})

		t.Run("PaymentsService Get returns payment if the input was its ID", func(t *testing.T) {
			repo := new(mocks.PaymentsRepository)
			repo.On("Get", validPayment.ID).Return(validPayment, nil)
			ps := NewPaymentsService(repo)

			retPayment, err := ps.Get(validPayment.ID)

			assert.Nil(err)
			assert.Equal(validPayment, retPayment, "Retrieved payment differs from expected payment after Get")
			repo.AssertExpectations(t)
		})
	})

	t.Run("Delete payment", func(t *testing.T) {
		t.Run("Error if invalid ID", func(t *testing.T) {
			ps := NewPaymentsService(nil)

			err := ps.Delete("")

			assert.Error(err)
		})

		t.Run("Error if not exists", func(t *testing.T) {
			repo := new(mocks.PaymentsRepository)
			id := "non-existing-id"
			repo.On("Exists", id).Return(false)
			ps := NewPaymentsService(repo)

			err := ps.Delete(id)

			assert.Error(err)
			repo.AssertExpectations(t)
		})

		t.Run("Deletes if input valid", func(t *testing.T) {
			repo := new(mocks.PaymentsRepository)
			id := "existing-id"
			repo.On("Exists", id).Return(true)
			repo.On("Delete", id).Return(nil)
			ps := NewPaymentsService(repo)

			err := ps.Delete(id)

			assert.Empty(err)
			repo.AssertExpectations(t)
		})
	})
}
