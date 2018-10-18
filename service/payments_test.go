package service

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
	"github.com/mysza/paymentsapi/domain"
	"github.com/mysza/paymentsapi/service/mocks"
	"github.com/stretchr/testify/assert"
)

func loadBytes(t *testing.T, name string) []byte {
	path := filepath.Join("testdata", name)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return bytes
}

func paymentFromFile(t *testing.T, name string) *domain.Payment {
	bytes := loadBytes(t, name)
	payment, err := domain.PaymentFromByteSlice(bytes)
	if err != nil {
		t.Fatal(err)
	}
	return payment
}

func TestAdd(t *testing.T) {
	assert := assert.New(t)

	t.Run("PaymentsService Add returns error if invalid input passed", func(t *testing.T) {
		ps := NewPaymentsService(nil)
		payment := paymentFromFile(t, "invalidPayment.json")

		_, err := ps.Add(payment)

		assert.Error(err)
	})

	t.Run("PaymentsService Add returns added payment ID if the input was valid", func(t *testing.T) {
		payment := paymentFromFile(t, "validPayment.json")
		payment.ID = ""
		repo := new(mocks.PaymentsRepository)
		repo.On("Add", payment).Return("4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43", nil)
		ps := NewPaymentsService(repo)

		id, err := ps.Add(payment)

		assert.Equal("4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43", id, "Returned ID should equal expected value")
		assert.Nil(err)
		repo.AssertExpectations(t)
	})

	t.Run("PaymentsService Add returns error if payment ID was set", func(t *testing.T) {
		payment := paymentFromFile(t, "validPayment.json")
		ps := NewPaymentsService(nil)

		id, err := ps.Add(payment)

		assert.Empty(id, "Returned ID should be empty")
		assert.Error(err, "Error should be set")
	})
}

func TestGetAllReturnsAllPaymentsFromRepo(t *testing.T) {
	assert := assert.New(t)
	payment := paymentFromFile(t, "validPayment.json")
	payments := []*domain.Payment{payment, payment}
	repo := new(mocks.PaymentsRepository)
	repo.On("GetAll").Return(payments, nil)
	ps := NewPaymentsService(repo)

	retPayments, err := ps.GetAll()

	assert.Equal(payments, retPayments, "Returned payments should be all in repository")
	assert.Nil(err)
	repo.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	assert := assert.New(t)
	validPayment := paymentFromFile(t, "validPayment.json")
	invalidPayment := paymentFromFile(t, "invalidPayment.json")

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
}

func TestGet(t *testing.T) {
	ps := NewPaymentsService(nil)
	assert := assert.New(t)

	t.Run("PaymentsService Get returns error if invalid ID passed", func(t *testing.T) {
		retPayment, err := ps.Get("")

		assert.Error(err)
		assert.Empty(retPayment)
	})

	t.Run("PaymentsService Get returns payment if the input was its ID", func(t *testing.T) {
		validPayment := paymentFromFile(t, "validPayment.json")
		repo := new(mocks.PaymentsRepository)
		repo.On("Get", validPayment.ID).Return(validPayment, nil)
		ps := NewPaymentsService(repo)

		retPayment, err := ps.Get(validPayment.ID)

		assert.Nil(err)
		assert.Equal(validPayment, retPayment, "Retrieved payment differs from expected payment after Get")
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	assert := assert.New(t)

	t.Run("Error if invalid ID", func(t *testing.T) {
		ps := NewPaymentsService(nil)

		err := ps.Delete("")

		assert.Error(err)
	})

	t.Run("Error if not exists", func(t *testing.T) {
		repo := new(mocks.PaymentsRepository)
		id := uuid.New().String()
		repo.On("Exists", id).Return(false)
		ps := NewPaymentsService(repo)

		err := ps.Delete(id)

		assert.Error(err)
		repo.AssertExpectations(t)
	})

	t.Run("Deletes if input valid", func(t *testing.T) {
		repo := new(mocks.PaymentsRepository)
		id := uuid.New().String()
		repo.On("Exists", id).Return(true)
		repo.On("Delete", id).Return(nil)
		ps := NewPaymentsService(repo)

		err := ps.Delete(id)

		assert.Empty(err)
		repo.AssertExpectations(t)
	})
}
