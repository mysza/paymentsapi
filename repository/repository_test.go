package repository

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/dgraph-io/badger"
	"github.com/mysza/paymentsapi/test"
	"github.com/stretchr/testify/assert"
)

func createDB() (*badger.DB, string) {
	opts := badger.DefaultOptions
	dir, _ := ioutil.TempDir("", "paymentsapibadger")
	opts.Dir = dir
	opts.ValueDir = dir
	db, err := badger.Open(opts)
	if err != nil {
		panic(err)
	}
	return db, dir
}

func prepareRepository() (*PaymentsRepository, func()) {
	db, dir := createDB()
	cleanup := func() {
		db.Close()
		os.RemoveAll(dir)
	}
	return New(db), cleanup
}

func TestRepository(t *testing.T) {
	assert := assert.New(t)
	testDataDir := filepath.Join("..", "testdata")
	validPaymentNoID := test.PaymentFromFile(t, filepath.Join(testDataDir, "validPayment.json"))
	validPaymentNoID.ID = ""

	t.Run("Repository add", func(t *testing.T) {
		repo, cleanup := prepareRepository()
		defer cleanup()

		id, err := repo.Add(validPaymentNoID)

		assert.Nilf(err, "Error adding to repo: %v", err)
		assert.NotEmpty(id, "No ID was created")
	})

	t.Run("Repository get", func(t *testing.T) {
		repo, cleanup := prepareRepository()
		defer cleanup()

		id, _ := repo.Add(validPaymentNoID)

		retPayment, err := repo.Get(id)

		assert.Nilf(err, "Error getting from repo: %v", err)
		assert.NotEmpty(retPayment, "No payment returned")
		assert.Equalf(id, retPayment.ID, "Retruned payment ID differs from ID defined when adding")
	})

	t.Run("Repository get all", func(t *testing.T) {
		repo, cleanup := prepareRepository()
		defer cleanup()

		// create 10 instances of payments
		for ix := 0; ix < 10; ix++ {
			repo.Add(validPaymentNoID)
		}

		payments, err := repo.GetAll()

		assert.Nilf(err, "Error getting all from repo: %v", err)
		assert.NotEmpty(payments, "No payments returned")
		assert.Equalf(10, len(payments), "Number of payments is incorrect; expected: %v, actual: %v", 10, len(payments))
	})

	t.Run("Repository delete", func(t *testing.T) {
		repo, cleanup := prepareRepository()
		defer cleanup()

		id, _ := repo.Add(validPaymentNoID)

		err := repo.Delete(id)

		assert.Nilf(err, "Error deleting payment: %v", err)
		assert.False(repo.Exists(id), "Repo reports payment still exists")
	})

	t.Run("Repository update", func(t *testing.T) {
		repo, cleanup := prepareRepository()
		defer cleanup()

		id, _ := repo.Add(validPaymentNoID)
		payment, _ := repo.Get(id)
		orgID := "ACME Inc."
		payment.OrganisationID = orgID
		err := repo.Update(payment)
		updated, _ := repo.Get(id)

		assert.Nilf(err, "Error updating payment: %v", err)
		assert.Equalf(orgID, updated.OrganisationID, "Payment not updated; expected: %v, got: %v", orgID, updated.OrganisationID)
	})
}
