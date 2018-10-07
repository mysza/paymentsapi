package repository

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/dgraph-io/badger"
	"github.com/mysza/paymentsapi/domain"
	"github.com/stretchr/testify/assert"
)

func createDB() (*badger.DB, string) {
	opts := badger.DefaultOptions
	dir, _ := ioutil.TempDir("", "badger")
	opts.Dir = dir
	opts.ValueDir = dir
	db, err := badger.Open(opts)
	if err != nil {
		panic(err)
	}
	return db, dir
}

func TestAdd(t *testing.T) {
	db, dir := createDB()
	defer func() {
		// cleanup
		db.Close()
		os.RemoveAll(dir)
	}()
	repo := New(db)
	payment := domain.CreateTestPayment()
	payment.ID = ""
	assert := assert.New(t)
	t.Run("Add retuns ID of the created payment", func(t *testing.T) {
		id, err := repo.Add(payment)
		payment.ID = id
		assert.Nilf(err, "Error adding to repo: %v", err)
		assert.NotEmpty(id, "No ID was created")
	})
}

func TestGet(t *testing.T) {
	db, dir := createDB()
	defer func() {
		// cleanup
		db.Close()
		os.RemoveAll(dir)
	}()
	repo := New(db)
	payment := domain.CreateTestPayment()
	id, _ := repo.Add(payment)
	payment.ID = id
	assert := assert.New(t)
	t.Run("Get returns previously added payment", func(t *testing.T) {
		retPayment, err := repo.Get(id)
		assert.Nilf(err, "Error getting from repo: %v", err)
		assert.NotEmpty(retPayment, "No payment returned")
		assert.Equalf(payment, retPayment, "Retruned payment differs from added")
	})
}

func TestGetAll(t *testing.T) {
	db, dir := createDB()
	defer func() {
		// cleanup
		db.Close()
		os.RemoveAll(dir)
	}()
	repo := New(db)
	payment := domain.CreateTestPayment()
	// create 10 instances of payments
	for ix := 0; ix < 10; ix++ {
		repo.Add(payment)
	}
	assert := assert.New(t)
	t.Run("GetAll returns all payments", func(t *testing.T) {
		payments, err := repo.GetAll()
		assert.Nilf(err, "Error getting all from repo: %v", err)
		assert.NotEmpty(payments, "No payments returned")
		assert.Equalf(10, len(payments), "Number of payments is incorrect; expected: %v, actual: %v", 10, len(payments))
	})
}

func TestDelete(t *testing.T) {
	db, dir := createDB()
	defer func() {
		// cleanup
		db.Close()
		os.RemoveAll(dir)
	}()
	repo := New(db)
	payment := domain.CreateTestPayment()
	id, _ := repo.Add(payment)
	assert := assert.New(t)
	t.Run("Delete deletes the payment", func(t *testing.T) {
		err := repo.Delete(id)
		assert.Nilf(err, "Error deleting payment: %v", err)
		assert.False(repo.Exists(id), "Repo reports payment still exists")
	})
}

func TestUpdate(t *testing.T) {
	db, dir := createDB()
	defer func() {
		// cleanup
		db.Close()
		os.RemoveAll(dir)
	}()
	repo := New(db)
	payment := domain.CreateTestPayment()
	var id string
	// create 10 instances of payments
	for ix := 0; ix < 10; ix++ {
		id, _ = repo.Add(payment)
	}
	assert := assert.New(t)
	t.Run("Update updates the payment", func(t *testing.T) {
		payment, _ := repo.Get(id)
		orgID := "ACME Inc."
		payment.OrganisationID = orgID
		err := repo.Update(payment)
		updated, _ := repo.Get(id)
		assert.Nilf(err, "Error updating payment: %v", err)
		assert.Equalf(orgID, updated.OrganisationID, "Payment not updated; expected: %v, got: %v", orgID, updated.OrganisationID)
	})
}
